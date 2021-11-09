package foo

import (
	"strconv"
)

func String(n int) string {
	if n == 3 {
		return "Foo"
	} else if n == 5 {
		return "Bar"
	}
	return strconv.Itoa(n)
}
