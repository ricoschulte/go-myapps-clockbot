# myapps-clockbot

Demo App to use [go-myapps](https://github.com/ricoschulte/go-myapps).

## Build

``` BASH
go build -o go-myapps-clockbot .
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
```