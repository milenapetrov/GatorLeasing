package handler

import (
	"bytes"
	"encoding/json"
	"log"
)

func createBody(request interface{}) *bytes.Reader {
	data, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	return reader
}
