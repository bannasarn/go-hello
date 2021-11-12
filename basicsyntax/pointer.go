package basicsyntax

import "fmt"

func Pointer() {
	var x int = 69
	var px *int = &x

	fmt.Printf("value x = %v\n", x)
	fmt.Printf("address x = %v\n", &x)

	fmt.Printf("value y = %v\n", *px)
	fmt.Printf("address y = %v\n", px)

	*px = 0

	fmt.Printf("value x = %v\n", x)
	fmt.Printf("address x = %v\n", &x)

	fmt.Printf("value y = %v\n", *px)
	fmt.Printf("address y = %v\n", px)
}
