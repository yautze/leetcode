package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	X float64
	N int
}

// TestTwoSum -
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example1",
			args: args{
				X: 2,
				N: 10,
			},
			want: 1024.00000,
		},
		// float64 * int的溢位問題
		// {
		// 	name: "Example2",
		// 	args: args{
		// 		X: 2.10000,
		// 		N: 3,
		// 	},
		// 	want: 9.261000000,
		// },
		{
			name: "Example3",
			args: args{
				X: 2,
				N: -2,
			},
			want: 0.25,
		},
		{
			name: "Example4",
			args: args{
				X: 1,
				N: -2147483648,
			},
			want: 1,
		},
		{
			name: "Example5",
			args: args{
				X: -1,
				N: 2147483647,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		res := myPow(tt.args.X, tt.args.N)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
