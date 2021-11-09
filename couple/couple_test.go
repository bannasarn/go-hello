package couple_test

import (
	"testing"

	"github.com/bannasarn/go-hello/couple"
	"github.com/stretchr/testify/assert"
)

func TestCoupleGivenAWantsAstar(t *testing.T) {
	given := "A"
	want := []string{"A*"}

	get := couple.Slice(given)

	assert.Equal(t, get, want)
}

func TestCoupleGivenABCDEFWantsABCDEF(t *testing.T) {
	given := "ABCDEF"
	want := []string{"AB", "CD", "EF"}

	get := couple.Slice(given)

	assert.Equal(t, get, want)
}

func TestCoupleGivenABCDEFGWantsABCDEFGstar(t *testing.T) {
	given := "ABCDEFG"
	want := []string{"AB", "CD", "EF", "G*"}

	get := couple.Slice(given)

	assert.Equal(t, get, want)
}
