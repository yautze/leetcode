package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	Target   int
	Position []int
	Speed    []int
}

// TestCarFleets -
func TestCarFleets(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				Target:   12,
				Position: []int{10, 8, 0, 5, 3},
				Speed:    []int{2, 4, 1, 1, 3},
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				Target:   10,
				Position: []int{3},
				Speed:    []int{3},
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				Target:   100,
				Position: []int{0, 2, 4},
				Speed:    []int{4, 2, 1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		res := carFleet(tt.args.Target, tt.args.Position, tt.args.Speed)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
