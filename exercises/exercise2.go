package main

func Power(base, exponent int) int {
	result := 0
	for i := 0; i < exponent; i++ {
		if i == 0 {
			result = base
		} else {
			result = result * base
		}
	}
	return result
}
