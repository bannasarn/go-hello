package basicsyntax

func ForLoop() {
	values := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(values); i++ {
		println(values[i])
	}
}

func WhileLoop() {
	values := []int{1, 2, 3, 4, 5, 6}

	i := 0
	for i < len(values) {
		println(values[i])
		i++
	}
}

func ForRange() {
	values := []int{10, 20, 30, 40, 50, 60}

	// In case need to use index
	for i, v := range values {
		println(i, v)
	}

	// In case no need to use index
	for _, v := range values {
		println(v)
	}
}
