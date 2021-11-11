package basicsyntax

import "fmt"

func Pointer() {
	var p *string
	var s string = "xxx"

	p = &s
	fmt.Println(s, p, &s)

	*p = "yyy"
	fmt.Println(s, p, &s)
}
