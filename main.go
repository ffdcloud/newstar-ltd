package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Struct to store form data
type RegistrationData struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// Handler to serve the form
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form template
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		http.Error(w, "Error loading form", http.StatusInternalServerError)
		return
	}

	// Serve the form
	tmpl.Execute(w, nil)
}

// Handler to process form submission
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	print("Someone called this endpoint")

	// Capture submitted data
	data := RegistrationData{
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	// Display the registration details as response (for testing purposes)
	fmt.Fprintf(w, "Registration successful!\n\nDetails:\nFirst Name: %s\nLast Name: %s\nEmail: %s", data.FirstName, data.LastName, data.Email)
}

func main() {
	// Serve static form
	http.HandleFunc("/", formHandler)
	// Handle form submission
	http.HandleFunc("/submit", submitHandler)

	// Start server on port 80
	log.Println("Starting server on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
