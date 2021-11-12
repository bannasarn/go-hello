package basicsyntax

import (
	"fmt"
	"strconv"
)

type Int int

func (i Int) ToString() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}

func TypeCreate() {
	var x Int = 50
	fmt.Printf("%q\n", x.ToString())
	x.Set(15)
	fmt.Printf("%q\n", x.ToString())
}
