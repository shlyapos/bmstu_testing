package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"skema/app"
	"skema/config"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := config.NewConfig()
	go app.InitAndRun(config)
}

func TestSchemaCreateSuccess(t *testing.T) {

	message := map[string]interface{}{
		"name":  "Uml class diagram",
		"owner": 1,
	}

	bytesRepresentation, errMarshal := json.Marshal(message)
	resp, errRequest := http.Post("http://localhost:5015/schema", "application/json", bytes.NewBuffer(bytesRepresentation))

	if errMarshal != nil {
		t.Error("Error marshal: ", errMarshal)
	}

	if errRequest != nil {
		t.Error("Error request: ", errRequest)
	}

	if errRequest != nil || resp.StatusCode != http.StatusOK {
		t.Error(resp)
	}
}

func BenchmarkTestSchemaCreate(b *testing.B) {
	t := new(testing.T)

	for i := 0; i < b.N; i++ {
		TestSchemaCreateSuccess(t)
	}
}
