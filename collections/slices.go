package collections

import "fmt"

func Slices() {

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 230)

	fmt.Printf("Scores -> %v, length -> %v\n", scores, len(scores))

	//Slice ranges

	cities := []string{
		"Durban",
		"Abuja",
		"Port novo",
		"Beijing",
		"Moscow",
		"Delhi",
		"Phoneix",
	}

	rangeOne := cities[1:3]
	fmt.Println(rangeOne)

	rangeTwo := cities[2:]
	fmt.Println(rangeOne, rangeTwo)

	rangeThree := cities[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Sage")
	fmt.Println(rangeOne)
}
