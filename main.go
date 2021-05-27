package main

import "fmt"

func main() {
	// ARRAYS
	// longhand: var arr [# of elements]data-type = [# of elements]data-type{data entries}
	// shorthand: arr := [# of elements]data-type{data entries}

	// var ages [3]int = [3]int{20, 25, 30}
	//ages := [3]int{20, 25, 30}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	fmt.Println(names, len(names))

	// SLICES - dynamic array
	// longhand: var slice = []data-type{data entries}
	// shorthand: slice := []data-type{data entries}
	var scores = []int{100, 50, 60}
	scores[2] = 25

	// APPEND - returns an array or slice with a new entry at the end, does not mutate
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// RANGES - gets the range of a slice or array, and stores it into a new slice

	// returns element position 1 up to position 3 minus 1
	rangeOne := names[1:3]
	// returns element position 2 and up
	rangeTwo := names[2:]
	// return element position 0 to 3 minus 2
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
