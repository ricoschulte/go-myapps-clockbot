package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/ricoschulte/go-myapps/connection"
	"github.com/ricoschulte/go-myapps/devicesapp"
	"github.com/ricoschulte/go-myapps/handler"
)

var host = flag.String("host", "", "the IP address or Hostname of the Pbx. Needs to be the HTTP/TLS port")
var username = flag.String("username", "", "the Username of the myApps account")
var password = flag.String("password", "", "the Password of the myApps account")
var useragent = flag.String("useragent", "Clockbot (go-myapps)", "the UserAgent showe in the list of sessions of the myApps account")
var insecureskipverify = flag.Bool("insecureskipverify", false, "skip verify TLS/SSL certificates")
var debug = flag.Bool("debug", false, "show debug output")

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

	accountConfig.Handler.AddHandler(&handler.HandleUpdateAppsInfo{})
	accountConfig.Handler.AddHandler(&handler.HandleUpdateAppsComplete{})
	accountConfig.Handler.AddHandler(&handler.HandleUpdateOwnPresence{})

	// add a handler for the devices app to myapps
	devicesapp := devicesapp.NewDevicesApp()
	accountConfig.Handler.AddHandler(devicesapp.Handler)

	go accountConfig.StartSession(&wg)

	wg.Wait()
	select {}
}
