package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertEquals := func(got string, want string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q, but wanted %q", got, want)
		}
	}

	t.Run("Given a letter, When Repeat(), it returns it five times.", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertEquals(got, want, t)
	})

	t.Run("Given a letter and a times parameter, When Repeat(), Then it outputs the letter the given time.", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"
		assertEquals(got, want, t)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}
