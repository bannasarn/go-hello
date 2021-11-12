package main

func main() {
	var i interface {
		String() string
	}
	println(i)
}
