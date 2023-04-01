package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Welcome struct {
	Sale string
	Time string
}

func main() {
	// Set up the welcome message with the current time
	welcome := Welcome{"Sale Begins Now", time.Now().Format(time.Stamp)}

	// Load the HTML template file
	template := template.Must(template.ParseFiles("template/template.html"))

	// Set up the HTTP server
	mux := http.NewServeMux()

	// Serve static files from the "static" directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the root URL
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method == http.MethodPost {
			// Update the welcome message with the new sale message from the form
			r.ParseForm()
			if sale := r.Form.Get("sale"); sale != "" {
				welcome.Sale = sale
			}
		}

		// Render the HTML template with the welcome message
		if err := template.ExecuteTemplate(w, "template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	server := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server started at http://localhost:8000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
