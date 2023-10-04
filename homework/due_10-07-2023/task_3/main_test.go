package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Remove(aPtr *[]int, index int) []int {
	return append((*aPtr)[:index], (*aPtr)[index+1:]...)
}

func AddOrRemove(aPtr *[]int, elem int) {
	indecies := []int{}
	for i := range *aPtr {
		if (*aPtr)[i] == elem {
			indecies = append(indecies, i)
		}
	}
	if len(indecies) == 0 {
		*aPtr = append(*aPtr, elem)
		return
	}

	for i := range indecies {
		*aPtr = Remove(aPtr, indecies[i]-i)
	}
}

func TestAddOrRemove(t *testing.T) {
	a := []int{1, 2, 3, 4}

	for i := 0; i < 20; i++ {
		AddOrRemove(&a, i)
	}

	expected := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	require.Equal(t, expected, a)
}

func TestRemoveFirst(t *testing.T) {
	a := []int{1, 2, 3}
	a = Remove(&a, 0)
	expected := []int{2, 3}
	require.Equal(t, expected, a)
}

func TestRemoveSecond(t *testing.T) {
	a := []int{1, 2, 3}
	a = Remove(&a, 1)
	expected := []int{1, 3}
	require.Equal(t, expected, a)
}
