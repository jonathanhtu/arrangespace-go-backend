package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetArrangement(w http.ResponseWriter, r *http.Request) {

}

func CreateArrangement(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	id += "hi"
}

func ExportArrangement(w http.ResponseWriter, r *http.Request) {

}
