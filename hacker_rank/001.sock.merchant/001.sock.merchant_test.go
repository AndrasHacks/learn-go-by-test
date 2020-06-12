package main

import "testing"

func TestSockMerchant(t *testing.T) {

	assertEquals := func(got int, want int, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("Got %d, but wanted %d.", got, want)
		}
	}

	t.Run("Given only a pack of socks, When SockMerchant(), Then tells 1 pairs",
		func(t *testing.T) {
			want := 1
			got := SockMerchant(2, []int{1, 1})
			assertEquals(got, want, t)
		})

	t.Run("Given no pairs, When SockMerchant(), Then tells 0 pairs",
		func(t *testing.T) {
			want := 0
			got := SockMerchant(2, []int{0, 1})
			assertEquals(got, want, t)
		})

	t.Run("Given the example, When SockMerchant(), Then tells correct answer.",
		func(t *testing.T) {
			want := 2
			got := SockMerchant(7, []int{1, 2, 1, 2, 1, 3, 2})
			assertEquals(got, want, t)
		})
}
