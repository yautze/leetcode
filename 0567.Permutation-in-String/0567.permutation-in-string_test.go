package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	S1 string
	S2 string
}

// TestCheckInClusion -
func TestCheckInClusion(t *testing.T) {
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				S1: "ab",
				S2: "eidbaooo",
			},
			want: true,
		},
		{
			name: "Example2",
			args: args{
				S1: "ab",
				S2: "eidboaoo",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		res := checkInclusion(tt.args.S1, tt.args.S2)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
