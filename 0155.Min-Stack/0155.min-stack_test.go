package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMinStack -
func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	// return -3
	check1 := minStack.GetMin()
	fmt.Println(check1)
	assert.Equal(t, -3, check1)

	minStack.Pop()
	// return 0
	check2 := minStack.Top() 	
	fmt.Println(check2)
	assert.Equal(t, 0, check2)

	// return -2
	check3 := minStack.GetMin()  
	fmt.Println(check3)
	assert.Equal(t, -2, check3)
}
