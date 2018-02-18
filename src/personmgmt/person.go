package personmgmt

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"google.golang.org/appengine/datastore"
)

func handlePerson(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		findPerson(w, req)
	case http.MethodDelete:
		deletePerson(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func findPerson(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	id := vars["personID"]
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ctx := appengine.NewContext(req)
	person, err := getPersonEntity(ctx, id)
	if err == datastore.ErrNoSuchEntity {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func deletePerson(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	id := vars["personID"]
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ctx := appengine.NewContext(req)
	err := deletePersonEntity(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePersons(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getPersons(w, req)
	case http.MethodPost:
		addPerson(w, req)
	}
}

func getPersons(w http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)

	persons, err := getPersonEntities(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(persons)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func addPerson(w http.ResponseWriter, req *http.Request) {

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	person := Person{}
	err = json.Unmarshal(b, &person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := appengine.NewContext(req)
	err = savePersonEntity(ctx, person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

