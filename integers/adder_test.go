package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("add 2 and 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func TestSubtract(t *testing.T) {
	t.Run("subtract 2 from 4", func(t *testing.T) {
		sum := Subtract(4, 2)
		expected := 2

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	got := ArraySum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
