package main

import "fmt"

var slc_1 []int

func main() {
	slc_2 := make([]int, 0, 4)

	slc_2 = append(slc_2, 1)
	slc_2[0] = 0
	slc_2 = append(slc_2, 2)
	slc_2 = append(slc_2, 3)

	fmt.Println(slc_1, slc_2)
	fmt.Printf("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))
	fmt.Printf("slc_2 len :%d cap:%d \n", len(slc_2), cap(slc_2))

	var slc_3 []int = []int{1, 2, 3, 4}
	fmt.Println("SLC3: ", slc_3)

	slc_4 := []int{5, 6}
	slc_3 = append(slc_3, slc_4...) //Birden fazla slice'ı birleştirme

	fmt.Println("SLC3: ", slc_3)

	matrix := [][]int{
		{1, 2},
		{2, 3, 4},
	}

	fmt.Println("Matrix Slice:", matrix)
	fmt.Println("Matrix [0][] len:", len(matrix[0]), "cap:", cap(matrix[0]))
	fmt.Println("Matrix [1][] len:", len(matrix[1]), "cap:", cap(matrix[1]))
}
