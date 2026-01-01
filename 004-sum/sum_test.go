package sum

import (
	"reflect"
	"slices"
	"testing"
)

func assertSum(t testing.TB, input []int, expected, actual int) {
	if actual != expected {
		t.Fatalf("expected %d given %v, got %d", expected, input, actual)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		expected := 15

		assertSum(t, numbers, expected, sum)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		sum := Sum(numbers)
		expected := 6

		assertSum(t, numbers, expected, sum)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should collect sum", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !slices.Equal(actual, expected) {
			t.Errorf("expected %v got %v", expected, actual)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should sum tails only", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9}, []int{5, 1, 2, 3}, []int{})
		expected := []int{2, 9, 6, 0}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v got %v", expected, actual)
		}
	})
}
