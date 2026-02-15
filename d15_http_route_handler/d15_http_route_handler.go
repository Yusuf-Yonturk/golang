package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Anasayfaya Hoş geldin!")
}

func YHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Yusuf Hoş geldin!")
}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Earth!")
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData := `{"message":"Pew pew"}`
	fmt.Fprintln(w, jsonData)
}

func main() {
	//define route
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/hi", HiHandler)
	http.HandleFunc("/pew", JsonHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/Yusuf", YHandler)
	//mux özel

	//port
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Println("HATA:", err)
		return
	}

}
