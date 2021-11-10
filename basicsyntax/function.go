package basicsyntax

func Function() {
	x := sum(10, 20)
	println(x)

	a, b, c := sum2(30, 40)
	println(a, b, c)

	_, d, _ := sum2(30, 40)
	println(d)
}

func sum(a, b int) int {
	return a + b
}

func sum2(a, b int) (int, string, bool) {
	return a + b, "Hello", true
}
