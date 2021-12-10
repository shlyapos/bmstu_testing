package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"skema/app/model"
	"skema/app/repo"
	"skema/app/service"
	"skema/util"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type SchemaController struct {
	service *service.SchemaService
}

type Schema struct {
	Title string `json:"title"`
	Descr string `json:"descr"`
	Owner int    `json:"owner"`
}

func NewSchemaController(db *gorm.DB) *SchemaController {
	ctrl := new(SchemaController)

	schemaRepo := repo.NewSchemaRepo(db)
	commentRepo := repo.NewCommentRepo(db)

	ctrl.service = service.NewSchemaService(schemaRepo, commentRepo)

	return ctrl
}

func InitSchemaController(r *mux.Router, db *gorm.DB) {
	ctrl := NewSchemaController(db)

	r.HandleFunc("/schema", ctrl.createSchema).Methods("POST")
	r.HandleFunc("/schema/{id}", ctrl.getSchemaById).Methods("GET")
}

func (ctrl *SchemaController) createSchema(res http.ResponseWriter, req *http.Request) {
	util.Logger.WriteInfoLog("Received a request", "SchemaController", "createSchema")

	schema := new(model.Schema)
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&schema)

	if err != nil {
		util.Logger.WriteInfoLog("Error request", "SchemaController", "createSchema")

		if errors.As(err, &unmarshalErr) {
			util.SetResponse(res, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			util.SetResponse(res, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	err = ctrl.service.CreateSchema(schema)

	if err != nil {
		util.SetResponse(res, "Bad Request "+err.Error(), http.StatusBadRequest)
	}

	util.Logger.WriteInfoLog("Success request", "SchemaController", "createSchema")
	util.SetResponse(res, "Success", http.StatusOK)
}

func (ctrl *SchemaController) getSchemaById(res http.ResponseWriter, req *http.Request) {
	util.Logger.WriteInfoLog("Received a request", "SchemaController", "getSchemaById")

	params := mux.Vars(req)
	id64, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		util.Logger.WriteInfoLog("Error convert", "SchemaController", "getSchemaById")
		util.SetResponse(res, "Bad Request "+err.Error(), http.StatusBadRequest)
	}

	id := uint(id64)
	schema, err := ctrl.service.TakeSchemaById(id)

	if err != nil {
		util.Logger.WriteInfoLog("Error request", "SchemaController", "getSchemaById")
		util.SetResponse(res, "Bad Request "+err.Error(), http.StatusBadRequest)
	}

	util.Logger.WriteInfoLog("Success request", "SchemaController", "getSchemaById")
	util.SetResonseData(res, schema, http.StatusOK)
}
