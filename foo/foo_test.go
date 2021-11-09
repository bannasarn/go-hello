package foo_test

import (
	"testing"

	"github.com/bannasarn/go-hello/foo"
)

func TestFooBarGiven1Wants1(t *testing.T) {
	given := 1
	want := "1"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven2Wants2(t *testing.T) {
	given := 2
	want := "2"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven3WantsFoo(t *testing.T) {
	given := 3
	want := "Foo"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven4Wants4(t *testing.T) {
	given := 4
	want := "4"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFooBarGiven5WantsBar(t *testing.T) {
	given := 5
	want := "Bar"

	get := foo.String(given)

	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}
