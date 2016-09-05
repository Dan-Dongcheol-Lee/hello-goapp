// +build !appengine

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Called main()")
	http.ListenAndServe(":8080", nil)
}
