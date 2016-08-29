// +build !appengine

package main

import (
    "net/http"
    "log"
)

func main() {
    log.Println("Called main()")
    http.ListenAndServe(":8080", nil)
}

