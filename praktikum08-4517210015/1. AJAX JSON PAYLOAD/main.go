package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := fmt.Sprintf(
			"Hello, my name is %s. I'm %d years old %s.",
			payload.Name,
			payload.Age,
			payload.Gender,
		)
		w.Write([]byte(message))
		return
	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/save", handleSave)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
