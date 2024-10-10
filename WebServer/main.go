package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestData struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `son:"age"`
}
type ResponseData struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	http.HandleFunc(`/api`, handleRequest)
	http.HandleFunc(`/status`, handleStatus)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	responseData := ResponseData{
		Message: "Server is running Successfully",
		Status:  "OK!",
	}
	w.Header().Set("Context-Type", "application.json")
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error parsing json data ", http.StatusInternalServerError)
		return
	}
	w.Write(responseJSON)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	//Parse body
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}
	//create response
	responseData := ResponseData{
		Message: fmt.Sprintf("Hello %x ! We have recieved your data. Your age is : %v and email id is: %v .", requestData.Name, requestData.Age, requestData.Email),
		Status:  "Success",
	}
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "error parsing jason data", http.StatusInternalServerError)
	}
	w.Write(responseJSON)
}
