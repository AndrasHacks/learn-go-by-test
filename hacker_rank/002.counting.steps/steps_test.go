package main

import "testing"

func TestSteps(t *testing.T) {
	assertEquals := func(want int, got int, t *testing.T) {
		t.Helper()
		if want != got {
			t.Errorf("Wanted %d, but got %d", want, got)
		}
	}

	t.Run("Given UUU, When CountSteps(), Then zero valleys counted", func(t *testing.T) {
		want := 0
		got := CountSteps(3, "UUU")
		assertEquals(want, got, t)
	})

	t.Run("Given D, When CountSteps(), Then one valleys counted", func(t *testing.T) {
		want := 1
		got := CountSteps(1, "D")
		assertEquals(want, got, t)
	})

	t.Run("Given DDUUDDUDUUUD, When CountSteps(), Then two valleys counted", func(t *testing.T) {
		want := 2
		got := CountSteps(12, "DDUUDDUDUUUD")
		assertEquals(want, got, t)
	})
}
