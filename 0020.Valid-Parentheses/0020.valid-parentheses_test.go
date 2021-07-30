package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsValid -
func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "Example1",
			args: "()",
			want: true,
		},
		{
			name: "Example2",
			args: "()[]{}",
			want: true,
		},
		{
			name: "Example3",
			args: "(]",
			want: false,
		},
		{
			name: "Example4",
			args: "([)]",
			want: false,
		},
		{
			name: "Example5",
			args: "{[]}",
			want: true,
		},
	}

	for _, tt := range tests {
		res := isValid(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
