package tests

import (
	"fmt"
	"strings"
)

func EmailTest() {

	emailAddress := "madebyshedrackgmail.com"

	if strings.Contains(emailAddress, "@") {
		fmt.Println("Welcome to Onboard base")
	} else {
		fmt.Println("Email address is not correct")
	}
}
