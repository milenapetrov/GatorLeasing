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

func createBody(payload interface{}) *bytes.Reader {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err.Error())
	}

	return bytes.NewReader(body)
}
