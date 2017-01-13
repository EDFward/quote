package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edfward/quote/models"
	"github.com/gorilla/mux"
)

func main() {
	models.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/quote", handleQuote)
	r.HandleFunc("/quotes", getQuotes).Methods("GET")
	r.HandleFunc("/quotes", insertQuote).Methods("POST")
	http.Handle("/", r)

	hostport := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Server running on %s\n", hostport)
	log.Fatal(http.ListenAndServe(hostport, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func handleQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Unimplemented")
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	user, platform := r.FormValue("user"), r.FormValue("platform")
	quotes, err := models.GetQuotes(user, platform)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func insertQuote(w http.ResponseWriter, r *http.Request) {
	user, platform := r.FormValue("user"), r.FormValue("platform")
	// Required fields.
	content, author := r.PostFormValue("content"), r.PostFormValue("author")
	// Optional fields.
	var sourcePtr *string
	var sectionPtr *string
	source, section := r.PostFormValue("source"), r.PostFormValue("section")
	// Handle empty optional fields.
	if source != "" {
		sourcePtr = &source
	}
	if section != "" {
		sectionPtr = &section
	}

	_, err := models.AddQuote(
		user, platform, content, author, sourcePtr, sectionPtr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return updated list of quotes.
	getQuotes(w, r)
}
