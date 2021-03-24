package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"9.TaskSW/data"
	"9.TaskSW/utils"
)

func setHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

//ReadData read quadratic equation arguments to local DB
func ReadData(writer http.ResponseWriter, request *http.Request) {
	setHeader(writer)
	log.Println("Read new arguments...")

	var newArguments data.QuadraticEquation
	err := json.NewDecoder(request.Body).Decode(&newArguments)

	if err != nil {
		msg := utils.ErrMsg{ErrMsg: "provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if len(data.DB) != 0 {
		data.DB = data.DB[:0]
	}

	data.DB = append(data.DB, &newArguments)

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(data.DB)

}

//SolveEquation return number of quadratic equation roots
func SolveEquation(writer http.ResponseWriter, request *http.Request) {
	setHeader(writer)
	log.Println("Calculate quadratic equation roots...")

	for _, v := range data.DB {
		if !v.AlreadyCalculated {
			utils.CalcNumOfRoots()
		}
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(data.DB)
}
