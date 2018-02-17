package hello

import (
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func HelloFoo(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello Foo!"))
}

func TraceFoo(w http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	DoTraceFoo(c, w, req)
}

func DoTraceFoo(c context.Context, w http.ResponseWriter, req *http.Request) {

	log.Infof(c, "Called trace-foo")

	client := urlfetch.Client(c)

	//Gets 'hellobar' service name
	hellobar, _ := appengine.ModuleHostname(c, "hellobar", "", "")
	if hellobar == "" {
		hellobar = "localhost:9090"
	}
	log.Infof(c, "hellobar hostname: %s", hellobar)

	resp, err := client.Get("http://" + hellobar + "/trace-bar")
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

func RootPath(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Service is running OK"))
}