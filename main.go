package main

import "fmt"

func main() {

	age := 34
	name := "Abe"
	score := 4.3

	// Println with variable outputs
	fmt.Println("My age is", age, "and my name is", name+".")

	// Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v.\n", age, name)
	fmt.Printf("my age is %q and my name is %q.\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %0.1f points!\n", score)

	// Sprintf (Save Formatted Strings)
	var str = fmt.Sprintf("my age is %v and my name is %v.\n", age, name)
	fmt.Println(str)
}
