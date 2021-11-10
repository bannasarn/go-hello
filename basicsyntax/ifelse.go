package basicsyntax

func IfElse() {
	point := 50
	if point >= 80 && point <= 100 {
		println("Grade A")
	} else if point >= 60 && point < 80 {
		println("Grade B")
	} else {
		println("Grade C")
	}
}
