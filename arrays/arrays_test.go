package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	assertEquals := func(got int, want int, array []int, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("Got %d, but wanted %d, from %v", got, want, array)
		}
	}

	t.Run("Given an array of number, when Sum(), then gets the sum of all numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertEquals(got, want, numbers, t)
	})

	t.Run("Given varying length input, When Sum(), Then returns the sum of the numebrs.", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertEquals(got, want, numbers, t)
	})
}
