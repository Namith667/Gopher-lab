package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to represent the JSON structure
type RequestData struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type ResponseData struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {

	// Start the server
	http.HandleFunc("/api", handleRequest)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

	// Create a response
	responseData := ResponseData{
		Message: fmt.Sprintf("Hello, %s! We received your data. You are %d years old and your email is %s.", requestData.Name,requestData.Age, requestData.Email),
		Status:  "Success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
