package foo_test

import (
	"testing"

	"github.com/bannasarn/go-hello/foo"
)

// func TestFooBar(t *testing.T) {
// 	type spec struct {
// 		given int
// 		want  string
// 	}
// 	specs := []spec{
// 		{given: 1, want: "1"},
// 		{given: 2, want: "2"},
// 		{given: 3, want: "Foo"},
// 		{given: 4, want: "4"},
// 		{given: 5, want: "Bar"},
// 	}
// }

func TestFooBarGiven1Want1(t *testing.T) {
	given := 1
	want := "1"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven2Want2(t *testing.T) {
	given := 2
	want := "2"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven3WantFoo(t *testing.T) {
	given := 3
	want := "Foo"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven4Want4(t *testing.T) {
	given := 4
	want := "4"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven5WantBar(t *testing.T) {
	given := 5
	want := "Bar"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven6WantFoo(t *testing.T) {
	given := 6
	want := "Foo"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven7Want7(t *testing.T) {
	given := 7
	want := "7"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven8Want8(t *testing.T) {
	given := 8
	want := "8"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven9WantFoo(t *testing.T) {
	given := 9
	want := "Foo"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven10WantBar(t *testing.T) {
	given := 10
	want := "Bar"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
