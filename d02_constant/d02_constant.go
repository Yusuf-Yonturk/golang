package main

const literal string = "var0"

const (
	literal_1 = "var1"
	literal_2 = "var2"
	literal_3 = "var3"
)

const (
	literal_4 = iota
	literal_5
	literal_6
)

func main() {
	println(literal, literal_1, literal_2, literal_3, literal_4, literal_5, literal_6)
}
