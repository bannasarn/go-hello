package basicsyntax

import "fmt"

func VariadicFunction() {
	primes := []int{1, 3, 5, 7}
	printSlice("hello", primes...)
}

func printSlice(s string, i ...int) {
	for _, v := range i {
		fmt.Println(s, v)
	}
}
