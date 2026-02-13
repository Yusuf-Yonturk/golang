package main

import "fmt"

var arr_1 [3]int
var arr_2 = [5]int{0, 1, 2, 3, 4}

func main() {
	arr_3 := make([]int, 3) // tek tek atama zorunlu, sağ tarafa liste şeklinde tanımlayamıyoruz altta.
	arr_3[1] = 1            // {1,2,3} olmuyor sağ taraf kendi basına farklı yabancı bir tip []int {1,2,3} yapmak zorundayız.

	fmt.Println(arr_1, arr_2, arr_3)
	fmt.Printf("arr_1 len:%d\n", len(arr_1))
	fmt.Printf("arr_2 len:%d\n", len(arr_2))
}
