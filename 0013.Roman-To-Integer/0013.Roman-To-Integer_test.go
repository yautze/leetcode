package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


// TestRomanToInt -
func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Example1",
			args: "III",
			want: 3,
		},
		{
			name: "Example2",
			args: "IV",
			want: 4,
		},
		{
			name: "Example3",
			args: "IX",
			want: 9,
		},
		{
			name: "Example4",
			args: "LVIII",
			want: 58,
		},
		{
			name: "Example5",
			args: "MCMXCIV",
			want: 1994,
		},
	}

	for _, tt := range tests {
		res := romanToInt(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
