package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"skema/app/repo"
	"skema/app/service"
	"skema/util"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	service *service.UserService
}

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func NewUserController(db *gorm.DB) *UserController {
	ctrl := new(UserController)

	repo := repo.NewUserRepo(db)
	ctrl.service = service.NewUserService(repo)

	return ctrl
}

func InitUserController(r *mux.Router, db *gorm.DB) {
	ctrl := NewUserController(db)

	r.HandleFunc("/user", ctrl.createUser).Methods("POST")
	r.HandleFunc("/user/{login}", ctrl.getUserByLogin).Methods("GET")
	r.HandleFunc("/user/{login}", ctrl.deleteUserByLogin).Methods("DELETE")
}

func (ctrl *UserController) createUser(res http.ResponseWriter, req *http.Request) {
	var user User
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&user)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			util.SetResponse(res, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			util.SetResponse(res, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	util.SetResponse(res, "Success", http.StatusOK)
}

func (ctrl *UserController) getUserByLogin(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	login := params["login"]

	fmt.Printf("This is /user/%v", login)
}

func (ctrl *UserController) deleteUserByLogin(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	login := params["login"]

	fmt.Printf("This is /user/%v", login)
}
