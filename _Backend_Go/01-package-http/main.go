package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", func(response http.ResponseWriter, request *http.Request) {
		log.Printf("%+v", request)

		if request.Method != http.MethodGet {
			http.Error(response, "This method doesn't support!", http.StatusMethodNotAllowed)
			return
		}

		message := map[string]string{
			"message": "Welcome welcome",
			"info":    "Kien Pham",
		}

		response.Header().Set("Content-type", "application/JSON")
		response.Header().Set("ABC", "DEF")

		/*data, err := json.Marshal(message)
		if err != nil {
			http.Error(response, "JSON error", http.StatusInternalServerError)
			return
		}
		response.Write(data)*/

		json.NewEncoder(response).Encode(message)
	})

	log.Println("Hello world")
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
