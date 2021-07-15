package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLengthOfLongestSubstring -
func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Example1",
			args: "abcabcbb",
			want: 3,
		},
		{
			name: "Example2",
			args: "bbbbb",
			want: 1,
		},
		{
			name: "Example3",
			args: "pwwkew",
			want: 3,
		},
		{
			name: "Example4",
			args: "",
			want: 0,
		},
	}

	for _, tt := range tests {
		res := lengthOfLongestSubstring(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
