package main

import (
    "net/http"

    "log"
    "github.com/Dan-Dongcheol-Lee/hello-goapp/hellofoo/api"
)

func init() {
    log.Println("Called init()")
    http.HandleFunc("/hello", api.Hello)
    http.HandleFunc("/trace-foo", api.TraceFoo)
}