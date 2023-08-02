package conditionals

import "fmt"

func Conditionals() {

	age := 40

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}

	addUsers("Josh", "luigi", "yoshi", "bowser")
}

func addUsers(users ...string) {
	if users[0] == "Josh" {
		fmt.Println("Hey welcome man you are first")
	} else if users[0] == "Yoshi" {
		fmt.Println("Hey welcome yoshi, you are second")
	} else {
		fmt.Println("No users found in index")
	}

	for index, value := range users {
		if index == 1 {
			fmt.Printf("Continuing at pos %v\n", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}

		fmt.Printf("the value at pos %v is %v\n", index, value)
	}
}
