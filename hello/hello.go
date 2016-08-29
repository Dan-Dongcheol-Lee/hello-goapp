package hello

import (
    "net/http"
    "google.golang.org/appengine/log"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "io/ioutil"
)

func HelloFoo(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello foo!"))
}

func TraceFromFoo(w http.ResponseWriter, req *http.Request) {
    ctx := appengine.NewContext(req)

    log.Infof(ctx, "Called trace-from-foo")

    client := urlfetch.Client(ctx)

    //resp, err := client.Get("http://localhost:8080/trace-to-bar")
    resp, err := client.Get("https://hellofoo-dot-pc-robusta-test-2.appspot.com/trace-to-bar")
    if err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }

    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()

    log.Infof(ctx, "Response body from bar: %s", body)

    w.Write([]byte("Called trace to bar!"))
}

func TraceToBar(w http.ResponseWriter, req *http.Request) {
    ctx := appengine.NewContext(req)

    log.Infof(ctx, "Called trace-to-bar")

    w.Write([]byte("Received from trace-from-foo!"))
}