package main

import (
	"reflect"
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

	t.Run("Given two slices with numbers entered, When SumAll(), Then a slice returned with the sums in it", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, but wanted %v", got, want)
		}
	})

	t.Run("Given two slices, When SumAllTails(), Then a slice with the sum of tails returned.", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, but wanted %v", got, want)
		}
	})
}
