package go_median_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		numbers  [][]float64
		expected float64
	}{
		{
			"one element",
			[][]float64{{1}},
			1,
		},
		{
			"two elements",
			[][]float64{{1, 2}},
			1.5,
		},
		{
			"three elements",
			[][]float64{{1, 2, 3}},
			2,
		},
		{
			"four elements",
			[][]float64{{1, 2, 3, 4}},
			2.5,
		},
		{
			"unordered elements",
			[][]float64{{4, 1, 2, 3}},
			2.5,
		},
		{
			"multiple slices",
			[][]float64{{4, 1, 2, 3}, {1}, {1, 2, 3}},
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, Median(test.numbers...))
		})
	}
}
