package collections

import "fmt"

func MyArrays() {

	var ages [3]int = [3]int{20, 25, 30}
	fmt.Printf("Ages %v\n", ages)

	var prices = [3]int{100, 125, 350}
	fmt.Printf("Prices -> %v, lengths -> %v\n", prices, len(prices))

	names := [4]string{
		"Yoshi",
		"Mario",
		"Peach",
		"Bowser",
	}

	fmt.Printf("Names %v length %v\n", names, len(names))

}
