package basicsyntax

import (
	"fmt"
)

// Note
// - Slice can append element in slice

func Slice() {
	x := []int{1, 2, 3}
	fmt.Println(x)

	y := append(x, 4) // Copy x and then create new address
	x = append(x, 4)  // Update existing address

	// Check cueent value
	fmt.Printf("This x = %#v\n", x)
	fmt.Printf("This y = %#v\n", y)

	// Check Lenght of Slice
	fmt.Printf("Length x = %#v\n", len(x))
	fmt.Printf("Length y = %#v\n", len(y))

}

func SliceExample2() {
	// Index   0   1   2   3   4   5   6   7   8
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Printf("This x = %#v\n", x)

	// Get value
	v := x[1]
	fmt.Printf("This v = %v\n", v)

	// Get Period of Slice
	y := x

	y = x[1:]
	fmt.Printf("This y = %v\n", y)

	y = x[2:6]
	fmt.Printf("This y = %v\n", y)

	y = x[:7]
	fmt.Printf("This y = %v\n", y)

	y = x[:]
	fmt.Printf("This y = %v\n", y)

}

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for y := 0; y < dy; y++ {
		var a []uint8
		for x := 0; x < dx; x++ {
			a = append(a, uint8(x*y))
		}
		result = append(result, a)
	}
	return result
}
