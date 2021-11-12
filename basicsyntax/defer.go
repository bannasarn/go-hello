package basicsyntax

import (
	"fmt"
	"log"
)

func Defer() {

	var n = 1

	println(n)
	defer println(n)
	n += 1
	println(n)
}

func Defer2() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	fmt.Println("Start")
	log.Panic("Yeh!!")
	fmt.Println("End")
}
