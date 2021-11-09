package foo

import "fmt"

func String(n int) string {
	if n == 3 {
		return "Foo"
	} else if n == 5 {
		return "Bar"
	}
	return fmt.Sprintf("%d", n)
}
