package main

import (
	"log"
	"net/http"
)

func main() {
    log.Println("server started")
	http.HandleFunc("/webhook", webhookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}