package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"./handler"

	"github.com/gorilla/mux"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GO World")
}

func hellotutorial(w io.Writer) {
	fmt.Fprint(w, "go tutorial\n")
}

func handleRequests() {
	//Init Router
	r := mux.NewRouter().StrictSlash(true) //gorilla used
	//Router Handler
	r.HandleFunc("/", helloworld).Methods("GET")
	r.HandleFunc("/users", handler.AllUsers).Methods("GET")
	r.HandleFunc("/user/{name}/{email}", handler.NewUser).Methods("POST")
	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{name}/{email}", handler.UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":1323", r))
}

func main() {
	hellotutorial(os.Stdout)

	handler.InitialMigration()
	handleRequests()

}
