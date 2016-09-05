package main

import (
    "net/http"

    "google.golang.org/appengine"
    "google.golang.org/appengine/log"
)

func init() {
    http.HandleFunc("/hello-bar", helloBar)
    http.HandleFunc("/trace-bar", traceBar)
}

func helloBar(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello Bar!"))
}

func traceBar(w http.ResponseWriter, req *http.Request) {
    c := appengine.NewContext(req)
    log.Infof(c, "Called trace-bar")
    w.Write([]byte("Trace called"))
}