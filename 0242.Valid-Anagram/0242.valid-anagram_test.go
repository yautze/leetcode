package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsAnagram -
func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s string
		t string
		want bool
	}{
		{
			name: "Example1",
			s: "anagram",
			t: "nagaram",
			want: true,
		},
		{
			name: "Example2",
			s: "rat",
			t: "car",
			want: false,
		},
	}

	for _, tt := range tests {
		res := isAnagram(tt.s, tt.t)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
