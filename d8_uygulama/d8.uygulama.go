package main

import (
	"errors"
	"fmt"
)

var db = map[string]map[string]string{
	"123": {"name": "Yusuf", "role": "Lead Architect"},
	"345": {
		"name": "",
		"role": "unassigned",
	},
}

func listUser() {
	for id, info := range db {
		fmt.Printf("ID: %s  Name: %s  Role: %s\n", id, info["name"], info["role"])
	}
}

func findUser(id string) (map[string]string, error) {
	if user, ok := db[id]; ok {
		return user, nil
	} else {
		return nil, errors.New("User record not found in database!")
	}
}

func deleteUser(id string) error {
	if _, ok := db[id]; !ok {
		return errors.New("cannot delete: user id not found")
	} else {
		delete(db, id)
		return nil
	}
}
func uptadeUser(id string, uptades map[string]string) error {
	if user, ok := db[id]; !ok {
		return fmt.Errorf("User id %s not found", id)
	} else {
		for key, value := range uptades {
			user[key] = value
		}
	}
	return nil
}

func main() {

	listUser()

	id := "123"

	changes := map[string]string{
		"name": "Yusuf", "role": "Senior Software Architect",
	}

	if user, err := findUser(id); err == nil {
		fmt.Printf("Success! User found:\n %s (%s)\n", user["name"], user["role"])
	} else {
		fmt.Println("Hata:", err)
	}

	if err := deleteUser("345"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Kullanici silindi")
	}
	listUser()

	if err := uptadeUser("123", changes); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Kullanici basariyla guncellendi!")
		listUser()
	}

}
