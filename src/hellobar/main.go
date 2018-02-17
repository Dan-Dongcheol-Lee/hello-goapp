// +build !appengine

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Called main()")
	registerHandlers()
	http.ListenAndServe(":8080", nil)
}

func registerHandlers() {
	http.HandleFunc("/hello-bar", helloBar)
	http.HandleFunc("/trace-bar", traceBar)
	http.HandleFunc("/", rootPath)
}