package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"9.TaskSW/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port        string
	grabPrefix  string = apiPrefix + "/grab"  //api/v1/item/
	solvePrefix string = apiPrefix + "/solve" //api/v1/items/
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	fmt.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	router.HandleFunc(grabPrefix, handlers.ReadData).Methods("POST")
	router.HandleFunc(solvePrefix, handlers.SolveEquation).Methods("GET")

	fmt.Println("Router created, start working")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
