package two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxNumLTN(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		n      int
		want   int
	}{
		{"Case 1", []int{1, 2, 3, 4, 5}, 231, 225},
		{"Case 2", []int{1, 3, 5, 7, 9}, 100, 99},
		{"Case 3", []int{0, 2, 4, 6, 8}, 753, 688},
		{"Case 4", []int{1, 2, 5, 8}, 542, 528},
		{"Case 5", []int{1, 2, 9, 4}, 2533, 2499},
		{"Case 6", []int{1, 2, 5, 4}, 2543, 2542},
		{"Case 7", []int{1, 2, 5, 4}, 2541, 2525},
		{"Case 8", []int{1, 2, 9, 4}, 2111, 1999},
		{"Case 9", []int{5, 9}, 5555, 999},
		{"Case 10", []int{3, 4, 5}, 2355, 555},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getMaxNumLtN(tt.digits, tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
