package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	patterns := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{10, -2, 8},
		{-10, -2, -12},
	}

	for idx, pattern := range patterns {
		actual := Add(pattern.a, pattern.b)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, actual)
		}
	}
}
