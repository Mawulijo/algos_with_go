package main

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestReserveString(t *testing.T) {
	is := is.New(t)

	is.Equal(ReserveStringR("hell, oh!"), "!ho ,lleh")
	is.Equal(ReserveStringB("hello"), "olleh")
	is.Equal(ReserveStringF("hello"), "olleh")
	is.Equal(ReserveStringSB("hello"), "olleh")
}