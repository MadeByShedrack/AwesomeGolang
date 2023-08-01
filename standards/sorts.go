package standards

import (
	"fmt"
	"sort"
)

func Sorts() {
	ages := []int{
		45,
		20,
		35,
		30,
		75,
		60,
		50,
		25,
	}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 3090)
	fmt.Println(index)

	names := []string{
		"Yoshi",
		"Mario",
		"Peach",
		"Bowser",
		"Luigi",
	}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Bowser"))
}
