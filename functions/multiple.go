package functions

import (
	"fmt"
	"strings"
)

func MultipleFunc() {
	// sayHello("Hello mr")
	getuserNames([]string{"Tim", "Daly", "Morrison"}, sayHello)
	firstName, secondName := getInitials("Tifa lockhart")
	fmt.Println(firstName, secondName)

	nameOne, nameTwo := getInitials("abel shedrack")
	fmt.Println(nameOne, nameTwo)
}

func sayHello(userMessage string) {
	fmt.Printf("Message -> %v\n", userMessage)
}

func getuserNames(userNames []string, users func(string)) {
	for _, myUsers := range userNames {
		users(myUsers)
	}
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
