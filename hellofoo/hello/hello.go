package hello

import (
    "net/http"
    "google.golang.org/appengine/log"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "io/ioutil"
    "io"
)

func Hello(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello Foo!"))
}

func TraceFoo(w http.ResponseWriter, req *http.Request) {
    c := appengine.NewContext(req)

    log.Infof(c, "Called trace-foo")

    client := urlfetch.Client(c)

    //Gets 'hellobar' service name
    hellobar, err := appengine.ModuleHostname(c, "hellobar", "", "")

    log.Infof(c, "hellobar hostname: %s", hellobar)

    resp, err := client.Get("https://" + hellobar + "/trace-bar")
    if err != nil {
        http.Error(w,
            http.StatusText(http.StatusInternalServerError),
            http.StatusInternalServerError)
        return
    }
    defer func() {
        io.Copy(ioutil.Discard, resp.Body)
        resp.Body.Close()
    }()

    w.Write([]byte("Called trace end!"))
}
