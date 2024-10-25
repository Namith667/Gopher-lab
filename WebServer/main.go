package main

import (
	"WebServer/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc(`/api`, handleRequest)
	err := godotenv.Load()
	port := os.Getenv("PORT")
	if err != nil {
		log.Fatalf("Cannot load")
	}
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid post method", http.StatusMethodNotAllowed)
		return
	}
	//Parse body
	var requestData models.RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}
	//create response
	responseData := models.ResponseData{
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
