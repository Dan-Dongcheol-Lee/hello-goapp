package hello_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/Dan-Dongcheol-Lee/hello-goapp/hellofoo/hello"
	"google.golang.org/appengine/aetest"
)

func TestHelloFoo_Ok(t *testing.T) {
	//Arrange: dev server running
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	//Create a new request. urlStr doesn't matter
	req, err := inst.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatalf("Failed to create a req: %v", err)
	}

	//Act
	w := httptest.NewRecorder()
	hello.HelloFoo(w, req)

	//Assert
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Got an error from response body: %v", err)
	}
	if string(b) != "Hello Foo!" {
		t.Errorf("Expected 'Hello Foo!' but got: %s", b)
	}
}
