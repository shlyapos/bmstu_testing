package test_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"skema/app"
	"skema/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := config.NewConfig()
	go app.InitAndRun(config)
}

var _ = Describe("Schema test", func() {
	message := map[string]interface{}{
		"name":  "Uml class diagram",
		"owner": 1,
	}

	Describe("Creating new schema", func() {
		Context("With user's role", func() {
			It("Should add successfully", func() {
				bytesRepresentation, errMarshal := json.Marshal(message)
				_, errRequest := http.Post("http://localhost:5015/schema", "application/json", bytes.NewBuffer(bytesRepresentation))

				Expect(errMarshal).To(BeNil())
				Expect(errRequest).To(BeNil())
			})
		})
	})
})
