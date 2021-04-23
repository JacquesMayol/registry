package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"registry/models"
	u "registry/utils"
	"strconv"
)

var GetNodes = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetNodes()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetNodeById = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10,64)

	if err != nil {
		fmt.Println(err)
		log.Fatal("Node identifier is not numeric string")
	}

	data := models.GetNode(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateNode = func(w http.ResponseWriter, r *http.Request) {

	node := &models.Node{}

	// Parses the request body
	err := r.ParseForm()
	if err != nil {
		u.Respond(w, u.Message(false, "Error while parse request body"))
	}

	err = json.NewDecoder(r.Body).Decode(node)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
	}

	resp := node.Create()
	u.Respond(w, resp)
}

var UpdateNode = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10,64)

	if err != nil {
		log.Fatal("Node identifier is not numeric string")
	}

	node := models.GetNode(id)

	if node == nil {
		log.Fatal("Node not exists")
		return
	}

	err = json.NewDecoder(r.Body).Decode(node)

	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := node.Update()
	u.Respond(w, resp)
}
