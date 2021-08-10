package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateParenthesis -
func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []string
	}{
		{
			name: "Example1",
			args: 3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "Example2",
			args: 1,
			want: []string{"()"},
		},
	}

	for _, tt := range tests {
		res := generateParenthesis(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
