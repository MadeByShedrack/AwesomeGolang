package variables

import "fmt"

func Variables() {
	var nameOne string = "Mario"
	fmt.Println(nameOne)

	var nameTwo = "Luigi"
	fmt.Println(nameTwo)
	fmt.Printf("Type -> %T, Value -> %v\n", nameTwo, nameTwo)

	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Bowser"
	fmt.Println(nameOne, nameThree)

	nameFour := "Yoshi"
	fmt.Printf("Name four -> %v\n", nameFour)

	//Integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//Bits and memory

	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	//Floats

	var scoreOne float32 = -1.5
	scoreOne = 25.98
	var scoreTwo float64 = 888994848575747383838
	fmt.Println(scoreOne, scoreTwo)

	scoreThree := 12.3
	fmt.Printf("Type -> %T, Value -> %v\n", scoreThree, scoreThree)
}
