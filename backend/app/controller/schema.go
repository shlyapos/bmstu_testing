package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"describe.me/app/service"
	"describe.me/util"
	"github.com/gorilla/mux"
)

type SchemaController struct {
	service *service.SchemaService
}

type Schema struct {
	Title string `json:"title"`
	Descr string `json:"descr"`
	Owner int    `json:"owner"`
}

func NewSchemaController() *SchemaController {
	ctrl := new(SchemaController)
	ctrl.service = service.NewSchemaService()

	return ctrl
}

func InitSchemaController(r *mux.Router) {
	ctrl := NewSchemaController()

	r.HandleFunc("/schema", ctrl.createSchema).Methods("POST")
	r.HandleFunc("/schema/{id}", ctrl.getSchemaById).Methods("GET")
}

func (ctrl *SchemaController) createSchema(res http.ResponseWriter, req *http.Request) {
	var schema Schema
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&schema)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			util.SetResponse(res, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			util.SetResponse(res, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	fmt.Println(schema)
	util.SetResponse(res, "Success", http.StatusOK)
}

func (ctrl *SchemaController) getSchemaById(res http.ResponseWriter, req *http.Request) {

}
