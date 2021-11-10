package basicsyntax

import "unicode/utf8"

func Len() {
	// Warning about Len, Do not use len for count character
	nameEN := "BANNASARN"
	nameTH := "บรรณสาร"

	// Bad
	println(len(nameEN))
	println(len(nameTH))

	// Good
	println(utf8.RuneCountInString(nameEN))
	println(utf8.RuneCountInString(nameTH))
}
