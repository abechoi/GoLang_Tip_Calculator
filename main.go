package main

import "fmt"

func main() {

	// STRINGS

	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yoshi"
	fmt.Println(nameFour)

	// INTEGERS

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// BITS & MEMORY
	// int8 - Range: -128 to 127
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// uint8 - unsigned int, Range: 0 to 255
	// var numThree uint16 = 256

	// FLOAT
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 7867369.3123
	scoreThree := 1.5

}
