package main

import (
    "net/http"

    "log"
    "github.com/Dan-Dongcheol-Lee/hello-goapp/hello"
)

func init() {
    log.Println("Called init()")
    http.HandleFunc("/hello-foo", hello.HelloFoo)
    http.HandleFunc("/trace-from-foo", hello.TraceFromFoo)
    http.HandleFunc("/trace-to-bar", hello.TraceToBar)
}