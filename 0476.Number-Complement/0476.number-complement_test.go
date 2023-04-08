package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFindComplement -
func TestFindComplement(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "Example1",
			args: 5,
			want: 2,
		},
		{
			name: "Example2",
			args: 1,
			want: 0,
		},
	}

	for _, tt := range tests {
		res := findComplement(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
