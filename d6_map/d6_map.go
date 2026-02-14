package main

import "fmt"

func main() {
	//Map
	//Anahtar - Değer || Key - Value
	//Hızlı veri erişimi sağlarlar.
	//Anahtarlar benzersizdir (unique)
	//Değerler herhangi bir veri tipi olabilir.

	rehber := map[string]int{
		"Yusuf":  20,
		"Marcus": 1,
	}

	cafe := map[string]float64{
		"Coffee":            120,
		"Basque Cheesecake": 220,
		"NON":               1,
	}
	fmt.Println(rehber)
	fmt.Printf("%#v", cafe)

	fmt.Println("\nCoffee price:", cafe["Coffee"])
	fmt.Println("Basque Cheesecake:", cafe["Basque Cheesecake"])
	delete(cafe, "NON")
	fmt.Println(cafe)

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

	fmt.Println(personas)
	fmt.Println("123 kullanici adina sahip personanın ismi:", personas["123"]["name"])

}
