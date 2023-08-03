package memory

import "fmt"

func Pointers() {
	fullName := "Dom Lucre"

	// updateFullName("")

	// fmt.Printf("Memory address of name is %v\n", &fullName)

	m := &fullName

	// fmt.Println("Memory address of m is: ", m)
	// fmt.Println("Value at memory address: ", *m)

	fmt.Println(fullName)
	updateFullName(m)
	fmt.Println(fullName)
}

func updateFullName(fullName *string) {
	*fullName = "Matt Walsh"
}
