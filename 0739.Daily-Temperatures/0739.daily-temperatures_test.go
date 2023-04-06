package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDailyTemperatures -
func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "Example1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "Example2",
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "Example3",
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
	}

	for _, tt := range tests {
		res := dailyTemperatures(tt.temperatures)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
