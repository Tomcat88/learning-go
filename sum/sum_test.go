package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum all the integers in an array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("found %d, but was expecting %d, numbers: %v", sum, expected, numbers)
		}
	})
	t.Run("should sum all the integers in a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("found %d, but was expecting %d, numbers: %v", sum, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	checkSums(t, got, expected)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})
	t.Run("empty array should yield 0", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2}, []int{3, 4, 5})
		expected := []int{0, 2, 9}
		checkSums(t, got, expected)

	})
}

func checkSums(t *testing.T, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, but expected %v", got, expected)
	}
}
