package leetcode

type MinStack struct {
	data       []int
	minHistory []int
}

func Constructor() MinStack {
	return MinStack{
		data:       make([]int, 0),
		minHistory: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	ms.data = append(ms.data, val)

	if len(ms.minHistory) < 1 || ms.minHistory[len(ms.minHistory)-1] > val {
		ms.minHistory = append(ms.minHistory, val)
	} else {
		ms.minHistory = append(ms.minHistory, ms.minHistory[len(ms.minHistory)-1])
	}
}

func (ms *MinStack) Pop() {
	ms.data = ms.data[:len(ms.data)-1]
	ms.minHistory = ms.minHistory[:len(ms.minHistory)-1]
}

func (ms *MinStack) Top() int {
	return ms.data[len(ms.data)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minHistory[len(ms.minHistory)-1]
}
