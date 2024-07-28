package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Go recieve message")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	textMessage := requestData["message"]

	if textMessage != "" {
		w.Write([]byte("Message: " + textMessage))
	} else {
		w.Write([]byte("please, input text message"))
	}
}

func main() {
	log.Print("Go server have stared")
	mux := http.NewServeMux()
	mux.HandleFunc("/message", MessageHandler)

	err := http.ListenAndServe(":6030", mux)
	log.Fatal(err)
}
