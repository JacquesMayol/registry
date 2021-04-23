package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"registry/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/nodes", controllers.GetNodes).Methods("GET")
	router.HandleFunc("/api/nodes/{id:[0-9]+}", controllers.GetNodeById).Methods("GET")
	router.HandleFunc("/api/nodes/create", controllers.CreateNode).Methods("POST")
	router.HandleFunc("/api/nodes/{id:[0-9]+}", controllers.UpdateNode).Methods("PATCH")

	// Attach JWT auth middleware
	// router.Use(app.JwtAuthentication)

	// router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // localhost
	}

	fmt.Println(port)

	// Launch the app, visit localhost:5000/api
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
