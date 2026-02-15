package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func Birthday(u *User) {
	u.Age++
}

func main() {
	number := 10
	fmt.Println(number)
	fmt.Println(&number)

	var ptr *int
	ptr = &number

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 20
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&ptr)
	fmt.Println(ptr)
	//Adres adresi tutuyor pointer to pointer

	usr1 := User{"Yusuf", 2}
	fmt.Println(usr1)
	Birthday(&usr1)
	fmt.Println(usr1)

}
