package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

type Address struct {
	City    string
	Country string
}

type Employee struct {
	UserInfo User
	Address  Address
	Salary   float64
}

func main() {
	user1 := User{
		Name:  "Yusuf",
		Age:   2,
		Email: "yusufyonturk.kisisel@gmail.com",
	}
	fmt.Println(user1)
	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)
	fmt.Println("Email:", user1.Email)
	var user2 User
	user2.Name = "Marcus"
	fmt.Println(user2)

	employee := Employee{
		UserInfo: user1,
		Address: Address{
			City:    "İstanbul",
			Country: "Turkey",
		},
		Salary: 240000,
	}
	fmt.Println(employee)

	fmt.Println("Employee Information:")
	fmt.Println("Name:", employee.UserInfo.Name)
	fmt.Println("Age:", employee.UserInfo.Age)
	fmt.Println("Email:", employee.UserInfo.Email)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("Country:", employee.Address.Country)
	fmt.Println("Salary:", employee.Salary)

	employees := []Employee{
		{
			UserInfo: User{"Yusuf", 2, "yusufyonturk.kisisel@gmail.com"},
			Address:  Address{"Istanbul", "Turkey"},
			Salary:   240000,
		},
		{
			UserInfo: User{"Marcus", 2, "xxx.xxx@gmail.com"},
			Address:  Address{"Istanbul", "Turkey"},
			Salary:   240000,
		},
	}
	for _, v := range employees {
		fmt.Println(v)
	}
}
