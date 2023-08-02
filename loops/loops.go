package loops

import "fmt"

func Loops() {
	fmt.Println("Welcome to loops")

	index := 0
	for index <= 250 {
		fmt.Printf("Hello world -> %v\n", index)
		index += 1
	}

	myIndex := 0
	sum := 0

	for myIndex <= 35 {
		sum += myIndex
		myIndex += 1
	}

	fmt.Printf("Sum -> %v\n", sum)

	for i := 0; i < 5; i++ {
		fmt.Printf("Index -> %v\n", i)
	}

	names := []string{
		"Mario",
		"Luigi",
		"Yoshi",
		"Peach",
	}

	for n := 0; n < len(names); n++ {
		fmt.Printf("Names -> %v\n", names[n])
	}

	for index, values := range names {
		fmt.Printf("the position at index %v is %v\n", index, values)
	}

	for _, name := range names {
		fmt.Printf("The names are %v\n", name)
		name = "Martha"
		fmt.Println(name)
	}

	fmt.Println(names)
}
