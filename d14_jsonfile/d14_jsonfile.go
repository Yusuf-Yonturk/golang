package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user1 := User{"Yusuf", 2, "yusufyonturk.kisisel@gmail.com"}

	data, err := json.MarshalIndent(user1, "", " ")
	if err != nil {
		fmt.Println("Json hatasi:", err)
		return
	}
	err = os.WriteFile("user.json", data, 0644)
	if err != nil {
		fmt.Println("Json hatasi:", err)
		return
	}
	fmt.Println("Json yazildi.")

	fileData, err := os.ReadFile("user.json")
	if err != nil {
		fmt.Println("Json hatasi:", err)
		return
	}
	var readUser User
	err = json.Unmarshal(fileData, &readUser)
	if err != nil {
		fmt.Println("Json hatasi:", err)
		return
	}
	fmt.Println(readUser)
	fmt.Println(readUser.Name)
	fmt.Println(readUser.Age)
	fmt.Println(readUser.Email)

}
