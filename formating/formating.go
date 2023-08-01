package formating

import "fmt"

func FormatingTypes() {

	//Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("New line \n")

	fmt.Println("Hello ninjas")

	age := 24
	name := "Shedrack"

	// fmt.Println("my age is ", age, "and my name is", name)

	fmt.Printf("My age is %v and my name is %v\n", age, name)

	countries := []string{
		"Japan",
		"Germany",
		"Austria",
		"Ukraine",
		"Niger",
	}

	countries = append(countries, "Denmark")
	fmt.Printf("Countries -> %T %v\n", countries, countries)

	for _, country := range countries {
		fmt.Printf("Users country -> %v\n", country)
	}

	fmt.Printf("My age is %q and my name is %q\n", age, name)

	fmt.Printf("Age is of type %T \n", age)

	fmt.Printf("You scored %0.1f points! \n", 225.55)

	greeting := "Hey good morning"
	greeter := "Tom welling"

	var statements = fmt.Sprintf("Greetings -> %v, Sender -> %v\n", greeting, greeter)
	fmt.Println("The statements is: ", statements)
}
