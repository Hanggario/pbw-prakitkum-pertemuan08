package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("Server started at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	// Data yang akan dikirim sebagai JSON
	data := []struct {
		Name string
		Age  int
	}{
		{"Garin Hanggario", 24},
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne", 21},
	}

	// Mengonversi data menjadi JSON
	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header response sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	// Mengirimkan data JSON sebagai respons
	_, err = w.Write(jsonInBytes)
	if err != nil {
		fmt.Println("Failed to write response:", err)
	}
}
