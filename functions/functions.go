package functions

import (
	"fmt"
	"math"
)

func Functions() {
	// sayGreetings("Hey good morning")
	// sayGreetings("Hallo, Guten tag")
	circleName([]string{"cloud", "tifa", "luigi"}, sayGreetings)

	areaOne := circleArea(10.5)
	areaTwo := circleArea(15)

	fmt.Println(areaOne, areaTwo)
	fmt.Printf("The circle 1 is of %0.3f and circle 2 is %0.3f\n", areaOne, areaTwo)
}

func sayGreetings(greeting string) {
	fmt.Printf("Greeting message %v\n", greeting)
}

func circleName(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}
