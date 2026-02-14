package main

import "fmt"

type Name string
type Mail string

func main() {

	var isim Name = "Yusuf"
	var email Mail = "yusufyonturk.kisisel@gmail.com\n"
	fmt.Println("Merhaba, ben", isim)
	fmt.Println(email, "\nYukaridaki epostadan bana ulasabilirsiniz")
}
