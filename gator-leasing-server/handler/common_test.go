package handler

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/faker"
)

func initializeTest() {
	faker.InitializeFaker()
}

func createBody(request interface{}) *bytes.Reader {
	data, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	return reader
}
