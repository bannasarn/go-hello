package main

import "fmt"

func main() {
	//println(swap(100, 200))
	//println(greeting("Bannasarn Boonmee"))

	//Array type Array
	// var array [5]string
	// primes := [...]int{2, 3, 5, 7, 11, 13}

	//Slice type Pointer
	var p []int
	primes := [...]int{1, 3, 5, 7}
	p = primes[:3]

	if p == nil {
		fmt.Println("It's nil")
	} else {
		fmt.Println(p)
	}

	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}

	fmt.Println(cap(p), len(p))

	p = append(p, 17, 19)
	fmt.Println(cap(p), len(p))
	fmt.Println(primes[3])
	fmt.Println(p[3])
}

// func isPrime(n int) bool {
// 	count := 0
// 	for i:=n; i>0; i-- {
// 		if n%1 == 0 {
// 			count ++
// 		}
// 		return count == 2
// 	}
// }

// Condition
// if ok := IsCorrect(); ok {
// 	println("In's correct")
// }
// func IsCorrect() bool {
// 	return true
// }

// Multiple return
func swap(a, b int) (int, int) {
	return b, a
}

// Hello World
func greeting(name string) string {
	return "Hello " + name + " (^3^)"
}
