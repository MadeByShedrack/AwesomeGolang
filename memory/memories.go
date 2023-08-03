package memory

import "fmt"

func GoMemory() {
	name := "Tifa"
	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"Pie":       5.99,
		"Ice cream": 3.99,
	}

	updateMenu(menu)

	fmt.Println(menu)
}

func updateName(x string) string {
	return x
}

func updateMenu(menu map[string]float64) {
	menu["coffee"] = 2.99
}
