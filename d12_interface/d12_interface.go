package main

//Bir davranış(method) imzasını tanımlar.
import "fmt"

type Greeter interface {
	Greet()
}
type RomanGreeter struct{}
type BritishGreeter struct{}

func (r RomanGreeter) Greet() {
	fmt.Println("Sastipe")
}

func (b BritishGreeter) Greet() {
	fmt.Println("Hey,mate!")
}

func SayHello(g Greeter) {
	g.Greet()
}

func main() {
	rom := RomanGreeter{}
	gb := BritishGreeter{}
	SayHello(rom)
	SayHello(gb)
}
