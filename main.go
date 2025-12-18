package main

import (
	"net/http"
	"fmt"
)

// HTTP Handler
func whatsappHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Data Collection
	// Body
	body := r.FormValue("Body")

	// Date (HTTP) Header
	date := r.Header.Get("Date")

	// MessageSid
	messageSid := r.FormValue("MessageSid")

	// Print Data
	fmt.Println("-------------------------------")
	fmt.Printf("Body: %s\n", body)
	fmt.Printf("Date: %s\n", date)
	fmt.Printf("MessageSid: %s\n", messageSid)
	fmt.Println("-------------------------------")

	// Respond with TwiML (XML)
	w.Header().Set("Content-Type", "text/xml")
	w.Write([]byte('<Response><Response>'))
}

// Start and Setup
func main() {
	http.HandleFunc("/whatsapp", whatsappHandler)
	http.ListenAndServe(":8080", nil)
}
