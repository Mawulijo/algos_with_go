package main

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestSumOfNumsInList(t *testing.T) {
	is := is.New(t)
	list := []int{1, 2, 3, 4, 5}
	emptylist := []int{}

	is.Equal(SumNumsInList(list), 15)
	is.Equal(SumNumsInList(emptylist), 0)
	is.Equal(SumNumsInListR(list), 15)
}