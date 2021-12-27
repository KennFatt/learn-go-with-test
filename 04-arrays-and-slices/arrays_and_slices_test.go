package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertSuccessMessage := func(t testing.TB, got, expect int) {
		t.Helper()

		if got != expect {
			t.Errorf("Expecting %q got %q", expect, got)
		}
	}

	t.Run("Sum array of numbers", func(t *testing.T) {
		numbers := []int{1, 1, 1, 1, 1}

		got := Sum(numbers)
		expect := 5

		assertSuccessMessage(t, got, expect)
	})

	t.Run("Sum slice of numbers", func(t *testing.T) {
		numbers := []int{1, 1, 1}

		got := Sum(numbers)
		expect := 3

		assertSuccessMessage(t, got, expect)
	})

}

// With len() array: 531488028                2.183 ns/op
// With len() slice: 490402198                2.360 ns/op
//
// With range array: 260006656                4.575 ns/op
//
// Conclusion: len() much way faster because it just a normal `for` iteration.
// 	from i to N. Where N = len(numbers)
//
// range in the other hand, is indeed slower. Why? because it genereates
// 	(index, value) even we do not use the index.
//
// TL;DR: collection[i] is much faster.
//
// array is slightliy faster than slice because it has fixed size.
func BenchmarkSum(t *testing.B) {

	numbers := []int{1, 1, 1, 1, 1}
	for i := 0; i < t.N; i++ {
		Sum(numbers)
	}
}

func TestSumAll(t *testing.T) {

	reflectMessage := func(t testing.TB, got, expect []int) {
		t.Helper()

		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Expecting %v got %v", expect, got)
		}
	}

	t.Run("Sum tails of a slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expect := []int{3, 9}

		reflectMessage(t, got, expect)
	})

	t.Run("Sum tails whose slice has empty element", func(t *testing.T) {
		got := SumAll([]int{}, []int{0, 9})
		expect := []int{0, 9}

		reflectMessage(t, got, expect)
	})
}

func BenchmarkSumAll(t *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{123, 1231, 6112, 61}

	for i := 0; i < t.N; i++ {
		SumAll(a, b)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails(
		[]int{},
		[]int{230, 4, 5, 1},
	)
	expect := []int{0, 10}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expecting %v got %v", expect, got)
	}
}

func BenchmarkSumAllTails(t *testing.B) {
	a := []int{123, 123145, 51, 12312, 123123, 5124, 213_213123, 213123}
	b := []int{4819458, 293, 23, 129831898, 219381289, 123}

	for i := 0; i < t.N; i++ {
		SumAllTails(a, b)
	}
}
