package main

import (
    "net/http"

    "log"
    "github.com/Dan-Dongcheol-Lee/hello-goapp/hellofoo/hello"
)

func init() {
    log.Println("Called init()")
    http.HandleFunc("/hello", hello.Hello)
    http.HandleFunc("/trace-foo", hello.TraceFoo)
}