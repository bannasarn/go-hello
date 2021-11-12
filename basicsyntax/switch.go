package basicsyntax

import "fmt"

func Switch() {
	n := 2
	switch {
	case n == 1:
		fmt.Printf("%v", n)
	case n == 2:
		fmt.Printf("%v", n)
	default:
		fmt.Printf("%v", n)
	}
}
