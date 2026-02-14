package main

import (
	"errors"
	"fmt"
)

func greet(name string) {
	fmt.Println("Merhaba", name)

}

// İsimlendirilmemiş fonksiyon
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Hata: Sifira bolme hatasi")
	}
	return a / b, nil

}

// İsimlendirilmiş fonksiyon
func triangle(base, height float64) (area float64, err error) {
	area = base * height / 2
	if base < 0 || height < 0 {
		err = errors.New("Hata: Taban veya Yukseklik 0'dan kucuk olamaz!")
		return 0, err
	}
	return area, nil
}

// Anonim fonksiyon
var multiply = func(a, b int) int {
	return a * b
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Closure Fonksiyon bir fonksiyon üreten bir fonksiyon durum state bilgisini saklayabilen fonksiyon tipi.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	greet("Yusuf")

	result, err := divide(10.0, 2)
	if err != nil {
		fmt.Printf("Eyvah bir hata cikti %v\n", err)
	} else {
		fmt.Printf("Sonuc: %f\n", result)
	}

	area, err := triangle(10, 2)
	if err != nil {
		fmt.Printf("Eyvah bir hata cikti %v \n", err)
	} else {
		fmt.Printf("Ucgenin alani: %f\n", area)
	}

	anonim := multiply(4, 3)
	fmt.Printf("Sonuc: %d\n", anonim)

	fmt.Printf("%d\n", factorial(5))

	sayac := makeCounter()
	sayac()
	fmt.Printf("%d\n", sayac())

}
