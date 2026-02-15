package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonData := `{"name":"Yusuf","age":2,"email":"yusufyonturk.kisisel@gmail.com"}`

	var user1 User
	err := json.Unmarshal([]byte(jsonData), &user1)
	if err != nil {
		fmt.Println("Bir hata var:", err)
		return
	}
	fmt.Println(user1.Name)
	fmt.Println(user1.Age)
	fmt.Println(user1.Email)

	newUser := User{
		Name:  "Marcus",
		Age:   1,
		Email: "xxxx.xxxx@xxxx.com",
	}

	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println("Json hatasi:", err)
	}
	fmt.Println(jsonBytes)
	fmt.Println(string(jsonBytes))
}
