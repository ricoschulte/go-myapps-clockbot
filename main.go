package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ricoschulte/go-myapps/connection"
)

var host = flag.String("host", "", "the IP address or Hostname of the Pbx. Needs to be the HTTP/TLS port")
var username = flag.String("username", "", "the Username of the myApps account")
var password = flag.String("password", "", "the Password of the myApps account")
var useragent = flag.String("useragent", "Clockbot (go-myapps)", "the UserAgent showe in the list of sessions of the myApps account")
var insecureskipverify = flag.Bool("insecureskipverify", false, "skip verify TLS/SSL certificates")
var debug = flag.Bool("debug", false, "show debug output")
var interval = flag.String("interval", "1s", "interval in seconds to update the user presence")
var timezone = flag.String("timezone", "Europe/Berlin", "the timezone to use to show the date & time")
var format = flag.String("format", "2006-01-02 15:04:05", "the format to show the current data & time")

func main() {

	flag.Usage = func() {
		fmt.Println("Usage: go-myapps-clockbot [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()
	options := ". use -h to show commandline options"
	if *host == "" {
		fmt.Println("host cant be empty", options)
		os.Exit(1)
	}

	if *username == "" {
		fmt.Println("username cant be empty", options)
		os.Exit(1)
	}

	if *password == "" {
		fmt.Println("password cant be empty", options)
		os.Exit(1)
	}
	if *useragent == "" {
		fmt.Println("useragent cant be empty", options)
		os.Exit(1)
	}

	_, err_location := time.LoadLocation(*timezone)
	if err_location != nil {
		fmt.Printf("unknown timezone '%s': %v\n", *timezone, err_location)
		os.Exit(1)
	}

	_, err_duration := time.ParseDuration(*interval)
	if err_duration != nil {
		fmt.Printf("unknown duration '%s': %v\n", *interval, err_duration)
		os.Exit(1)
	}

	var wg sync.WaitGroup

	accountConfig := connection.Config{
		Host:               *host,
		Username:           *username,
		Password:           *password,
		InsecureSkipVerify: *insecureskipverify,
		UserAgent:          *useragent,
		SessionFilePath:    fmt.Sprintf("myapps_session_%s.json", *username),
		Debug:              *debug,
	}

	// add a handler to the account
	// this handler will wait until after the login the pbx send a 'UpdateAppsComplete' message
	accountConfig.Handler.AddHandler(&ClockbotHandler{})

	go accountConfig.StartSession(&wg)

	wg.Wait()
	select {}
}

type ClockbotHandler struct{}

func (m *ClockbotHandler) GetMt() string {
	return "UpdateAppsComplete"
}

func (m *ClockbotHandler) HandleMessage(myAppsConnection *connection.MyAppsConnection, message []byte) error {
	// we received the UpdateAppsComplete message and we start the loop that sends the presence updates
	go m.RunLoop(myAppsConnection, message)
	return nil
}

func (m *ClockbotHandler) RunLoop(myAppsConnection *connection.MyAppsConnection, message []byte) error {
	for {
		tz, _ := time.LoadLocation(*timezone)
		// Get the current time in the specified timezone
		now := time.Now().In(tz)

		var message_out struct {
			Mt       string `json:"mt"`
			Activity string `json:"activity"`
			Note     string `json:"note"`
		}
		message_out.Mt = "SetOwnPresence"
		message_out.Activity = ""
		message_out.Note = GetFormatedTimeString(now, *format)
		updateStr, _ := json.Marshal(message_out)

		err_write := myAppsConnection.Conn.WriteMessage(websocket.TextMessage, []byte(updateStr))
		if err_write != nil {
			// if sending the update fails we quit our loop
			myAppsConnection.Config.Printf("error while sending message: %v", err_write)
			return err_write
		}
		// wait the interval and do it again
		duration, _ := time.ParseDuration(*interval)
		time.Sleep(duration)
	}

}

func GetFormatedTimeString(current time.Time, format string) string {
	return current.Format(format)
}
