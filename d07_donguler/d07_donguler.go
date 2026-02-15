package main

import "fmt"

func main() {
	personas := map[string]map[string]string{
		"123": {
			"name":  "Yusuf",
			"sname": "Marcus",
		},
		"456": {
			"name":  "Sebastian",
			"sname": "Helga",
		},
	}

	cafe := map[string]float64{
		"Coffee":            120,
		"Basque Cheesecake": 220,
		"NON":               1,
	}

	var arr_2 = [5]int{0, 1, 2, 3, 4}

	for i := 0; i < len(arr_2); i++ {
		fmt.Printf("Index:[%d]  =  %d\n", i, arr_2[i])
	}

	for index, value := range personas {
		fmt.Println("Index:", index, "Value:", value)
	}

	for key, value := range cafe {
		fmt.Println(key, "Fiyati:", value)
	}

	for key := range cafe {
		fmt.Println("Key:", key)
	}

	for _, value := range cafe {
		fmt.Println("Value:", value)
	}
}
