package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type WebhookRequest struct {
	Action		string
	Repository	struct {
		ID			string
		FullName	string
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var webhookData WebhookRequest
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("got webhook payload: ")
	fmt.Println(webhookData)
}