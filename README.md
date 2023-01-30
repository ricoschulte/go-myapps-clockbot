# go-myapps-clockbot

Demo App to use [go-myapps](https://github.com/ricoschulte/go-myapps).

Clockbot is a tool that helps you to set the presence of a user in a MyApps account to the current time. 
The Clockbot logs into a myApps account and sets the presence to the current time, based on the timezone, format, and update interval that are specified via CLI parameters.

With Clockbot, you can specify the timezone in which you want the current time to be displayed. 
You can choose from a list of valid IANA Time Zone Database names. 
This allows you to display the current time in your local time zone or any other time zone of your choice.

Additionally, Clockbot allows you to specify the format in which the current time is displayed. 
You can choose from a variety of time formats.
This ensures that the current time is displayed in a format that is easy for you to understand.

Finally, Clockbot also allows you to specify the update interval, which is the frequency at which the current time is updated in the myApps account.

## Build

``` BASH
go build -o go-myapps-clockbot .
```

Cross-compiling for a Raspberry Pi 4 or innovaphone App Platform on a innovaphone IP411

``` BASH
GOARCH=arm64 GOOS=linux go build . -o go-myapps-clockbot
```

## Usage

``` BASH
$ go-myapps-clockbot -h
Usage: go-myapps-clockbot [options]
Options:
  -debug
        show debug output
  -host string
        the IP address or Hostname of the Pbx. Needs to be the HTTP/TLS port
  -insecureskipverify
        skip verify TLS/SSL certificates
  -password string
        the Password of the myApps account
  -useragent string
        the UserAgent showe in the list of sessions of the myApps account (default "Clockbot (go-myapps)")
  -username string
        the Username of the myApps account
  -format string
        the format to show the current data & time (default "2006-01-02 15:04:05")  
  -timezone string
        the timezone to use to show the date & time (default "Europe/Berlin")
  -interval string
        interval in seconds to update the user presence (default "1s")
```

## Timezone

The `-timezone` option is used to set the timezone for the Clockbot. 
When this option is used, the Clockbot will display the current time in the specified timezone. 
The timezone can be specified as a string in the format of a valid IANA Time Zone Database name.

To find the list of valid timezone strings, you can refer to the IANA Time Zone Database (https://www.iana.org/time-zones).

Note: The -timezone option is case-sensitive, and must be written in the format as listed in the IANA Time Zone Database. 

If the specified timezone string is not valid, the Clockbot will return an error.

## Format

The `-format` option is used to set the format of the presence text. 

Here are some examples of format strings:

- **2006-01-02**: returns the date in the format `YYYY-MM-DD`
- **15:04:05**: returns the time in the format `HH:MM:SS`
- **2006-01-02 15:04:05**: returns the date and time in the format `YYYY-MM-DD HH:MM:SS`
- **Mon, 02 Jan 2006 15:04:05 GMT**: returns the date and time in the format `Mon, DD MMM YYYY HH:MM:SS GMT`
- **Monday, 02-Jan-06 15:04:05 PST**: returns the date and time in the format `Monday, DD-MMM-YY HH:MM:SS PST`

You can use any combination of these and other elements to create the desired format for your date and time. 
The complete reference for the format string can be found in the Go documentation: https://golang.org/src/time/format.go

If the specified format string is not valid, the Clockbot will return an error.

## Update interval

The `-interval` option is used to set the interval of updates. 
The string representation of a duration should be a sequence of decimal numbers, each with optional fraction and a unit suffix, such as `300ms`, `-1.5h` or `2h45m`. 
The recognized units are `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.

If the specified interval string is not valid, the Clockbot will return an error.

## About ©

[myApps](https://www.innovaphone.com/en/myapps/what-is-myapps.html) is a product of [innovaphone AG](https://www.innovaphone.com).

Documentation of the API used in this client can be found at [ innovaphone App SDK](https://sdk.innovaphone.com/).