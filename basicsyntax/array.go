package basicsyntax

import "fmt"

// Note
// - Array fixed length. So can not append

func Array() {
	// Define array with assign default value e.g. var x [3]int = [3]int{1, 2, 3}
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	fmt.Printf("%#v\n", x)
	// update array value
	x[0] = 10
	fmt.Println(x)

	// Array 2 dimension
	y := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(y)
	// update array value
	y[0][1] = 20
	fmt.Println(y)

	// Not define array capacity max
	z := [...]int{1, 2, 3, 4, 5}
	fmt.Println(z)
}
