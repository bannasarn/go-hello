package basicsyntax

import "fmt"

func Print(s string) {
	println("Hello " + s)
	fmt.Printf("Hello %v\n", s)
	fmt.Printf("Hello %#v\n", s)
}
