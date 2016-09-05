package main

import (
    "net/http"

    "log"
    "github.com/Dan-Dongcheol-Lee/hello-goapp/hellofoo/api"
    "github.com/mjibson/appstats"
)

func init() {
    log.Println("Called init()")
    http.HandleFunc("/hello-foo", api.HelloFoo)
    http.HandleFunc("/trace-foo", api.TraceFoo)
    http.HandleFunc("/trace-foo-stats", appstats.NewHandlerFunc(api.DoTraceFoo))
}