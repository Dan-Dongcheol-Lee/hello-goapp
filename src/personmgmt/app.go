package personmgmt

import (
	"net/http"
	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/persons/{personID}", handlePerson) 	// Find and Delete a Person
	r.HandleFunc("/persons/", handlePersons) 			// Get Persons and add a Person
	r.HandleFunc("/", rootPath)
	http.Handle("/", r)
}

func rootPath(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Service is running OK"))
}
