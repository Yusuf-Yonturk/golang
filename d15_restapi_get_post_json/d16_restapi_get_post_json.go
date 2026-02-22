package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
type CreateUserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		getUsers(w)
		return
	}
	if r.Method == http.MethodPost {
		createUser(w, r)
		return
	}
	http.Error(w, "Error: Only GET and POST method.", http.StatusMethodNotAllowed)

}

func getUsers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json") // Response JSON
	fileName := "users.json"

	info, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(fileName + "is not found, generating file...")

			empty := []User{}
			data, marshalErr := json.MarshalIndent(empty, "", " ")
			if marshalErr != nil {
				http.Error(w, "Could not create JSON"+marshalErr.Error(), http.StatusInternalServerError)
				return
			}

			writerErr := os.WriteFile(fileName, data, 0644)
			if writerErr != nil {
				http.Error(w, "Could not create file:"+writerErr.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "File Access Error"+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if info != nil {
			fmt.Println("File Name:", info.Name())
			fmt.Println("File Size:", info.Size())
			fmt.Println("Mod Time:", info.ModTime())
			fmt.Println("File Is Dir:", info.IsDir())
			fmt.Println("File Mode:", info.Mode())

		}
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}

	var users []User
	json.Unmarshal(data, &users)

	json.NewEncoder(w).Encode(users)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read body", http.StatusBadRequest)
		return
	}

	var newUser User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	if newUser.Name == "" {
		http.Error(w, "Name field required", http.StatusBadRequest)
		return
	}
	if newUser.Email == "" {
		http.Error(w, "Email field required", http.StatusBadRequest)
		return
	}
	if newUser.Age <= 0 {
		http.Error(w, "Age field required", http.StatusBadRequest)
		return
	}

	var users []User
	if _, err := os.Stat("users.json"); os.IsNotExist(err) {
		users = []User{}
		emptyFile, _ := json.MarshalIndent(users, "", "  ")
		os.WriteFile("users.json", emptyFile, 0644)
	} else {
		fileData, _ := os.ReadFile("users.json")
		json.Unmarshal(fileData, &users)
	}

	users = append(users, newUser)

	fileJson, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "Could not create JSON", http.StatusInternalServerError)
		return
	}

	os.WriteFile("users.json", fileJson, 0644)

	//API RESPONSE
	response := CreateUserResponse{
		Message: "User registered",
		User:    newUser,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/users", UsersHandler)

	http.ListenAndServe(":8080", nil)
}
