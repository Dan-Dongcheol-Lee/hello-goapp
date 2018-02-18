package personmgmt

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"fmt"
)

const personEntityName = "person-mgmt-Person"

type Person struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Address string `json:"address"`
}

func getPersonEntities(ctx context.Context) ([]Person, error) {

	query := datastore.NewQuery(personEntityName)
	it := query.Run(ctx)

	persons := []Person{}
	for {
		p := Person{}
		_, err := it.Next(&p)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get persons from datastore. error: %v", err)
		}
		persons = append(persons, p)
	}
	return persons, nil
}

func getPersonEntity(ctx context.Context, personID string) (Person, error) {
	person := Person{}
	k := datastore.NewKey(ctx, personEntityName, personID, 0, nil)
	err := datastore.Get(ctx, k, &person)
	return person, err
}


func savePersonEntity(ctx context.Context, person Person) error {
	key := person.ID
	k := datastore.NewKey(ctx, personEntityName, key, 0, nil)
	_, err := datastore.Put(ctx, k, &person)
	return err
}

func deletePersonEntity(ctx context.Context, personID string) error {
	k := datastore.NewKey(ctx, personEntityName, personID, 0, nil)
	return datastore.Delete(ctx, k)
}