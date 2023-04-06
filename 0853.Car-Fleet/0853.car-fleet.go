package leetcode

import "sort"

func carFleet(target int, position []int, speed []int) int {
	stackCarFleets := make([]float64, 0)
	// position mapping speed
	carMap := make(map[int]int)

	for i := 0; i < len(position); i++ {
		carMap[position[i]] = speed[i]
	}

	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j]
	})

	for _, v := range position {
		speed, _ := carMap[v]
		distance := target - v
		totalCostTime := float64(distance) / float64(speed)

		if len(stackCarFleets) > 0 {
			lastTime := stackCarFleets[len(stackCarFleets)-1]
			if totalCostTime > lastTime {
				stackCarFleets = append(stackCarFleets, totalCostTime)
			}
		} else {
			stackCarFleets = append(stackCarFleets, totalCostTime)
		}
	}

	return len(stackCarFleets)
}
