package main

import (
	"net/http"

	"log"

	"hellofoo/hello"
	"github.com/mjibson/appstats"
)

func init() {
	log.Println("Called init()")
	http.HandleFunc("/hello-foo", hello.HelloFoo)
	http.HandleFunc("/trace-foo", hello.TraceFoo)
	http.HandleFunc("/trace-foo-stats", appstats.NewHandlerFunc(hello.DoTraceFoo))
	http.HandleFunc("/", hello.RootPath)
}
