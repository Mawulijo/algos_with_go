package main

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestNumInList(t *testing.T) {
	is := is.New(t)
	list := []int{1, 2, 3, 4, 5}
	is.True(NumInList(list,5))
	is.False(NumInList(list,8))
}