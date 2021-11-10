package basicsyntax

import "fmt"

func VariadicFunction() {
	primes := []int{1, 3, 5, 7}
	printSlice(primes...)
}

func printSlice(i ...int) {
	for _, v := range i {
		fmt.Println(v)
	}
}
