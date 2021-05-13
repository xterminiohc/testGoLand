package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	userdb "sofka.com/mod/pkg/adapter"
	rest "sofka.com/mod/pkg/controller"
	auth "sofka.com/mod/pkg/controller/auth"
)

func main() {
	fmt.Println("Start Program")
	userdb.NewUser()
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", rootPage)
	myRouter.HandleFunc("/greet", auth.IsAuthorized(rest.Greet)).Methods("GET")
	myRouter.HandleFunc("/token", auth.GetToken).Methods("GET")
	http.Handle("/", myRouter)
	log.Fatal(http.ListenAndServe(":9091", nil))
	fmt.Println("server Start on: http://localhost:9091/")
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sofka Test Goland Restful\n ")
	fmt.Fprintf(w, "List of APIs:\n ")
	fmt.Fprintf(w, " 1) localhost:9091/greet   :: private Resource, send a Token for GET method.\n")
	fmt.Fprintf(w, " 2) localhost:9091/token   :: public resource to get a token.\n")
}
