package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	Str string
	Target string
}

// TestStrStr2 -
func TestStrStr2(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				Str: "hello",
				Target: "ll",
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				Str: "aaaaa",
				Target: "bba",
			},
			want: -1,
		},
		{
			name: "Example3",
			args: args{
				Str: "",
				Target: "",
			},
			want: 0,
		},
		{
			name: "Example4",
			args: args{
				Str: "mississippi",
				Target: "pi",
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		res := strStr2(tt.args.Str, tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
