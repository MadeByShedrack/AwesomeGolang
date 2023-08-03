package maps

import "fmt"

func MapFunc() {

	menu := map[string]float64{
		"Soup":          4.99,
		"Pie":           7.99,
		"Salad":         6.99,
		"Toffe pudding": 3.55,
	}

	for keys, value := range menu {
		fmt.Printf("keys -> %v, Values -> %v\n", keys, value)
	}

	fmt.Println(menu["Salad"])

	phoneBook := map[int]string{
		26889477: "Mario",
		67889345: "Luigia",
		34557689: "Tifa",
		87364645: "Axwell",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[87364645])

	phoneBook[67889345] = "Bowser"
	fmt.Println(phoneBook)

	phoneBook[87364645] = "Yoshi"
	fmt.Println(phoneBook)

	for contacts, names := range phoneBook {
		fmt.Printf("Contacts -> %v, Names -> %v\n", contacts, names)
	}
}
