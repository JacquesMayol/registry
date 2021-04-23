package controllers

import (
	"encoding/json"
	"net/http"
	"registry/models"
	u "registry/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	// Decode the request body into struct and failed if any error occur
	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	// Create account
	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	// Decode the request body into struct and failed if any error occur
	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
