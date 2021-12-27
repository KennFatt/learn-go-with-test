package iterations

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, expect string) {
		t.Helper()

		if got != expect {
			t.Errorf("Expecting %q got %q", expect, got)
		}
	}

	t.Run("repeat \"a\" five times", func(t *testing.T) {
		got := Repeat("a", 5)
		expect := "aaaaa"

		assertCorrectMessage(t, got, expect)
	})

	t.Run("repeat \"b\" N times", func(t *testing.T) {
		n := 3
		letter := "b"

		got := Repeat(letter, n)
		expect := "bbb"

		assertCorrectMessage(t, got, expect)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
