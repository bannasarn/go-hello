package main

import "github.com/bannasarn/go-hello/couple"

func main() {
	couple.Slice("ABCDEF")
}

// Multiple return
func swap(a, b int) (int, int) {
	return b, a
}

// Hello World
func greeting(name string) string {
	return "Hello " + name + " (^3^)"
}
