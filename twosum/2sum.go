package twosum

import (
	"sort"
)

type number struct {
	value int
	index int
}

func TwoSum(nums []int, target int) []int {
	indexedNums := make([]number, len(nums))
	for index, value := range nums {
		indexedNums[index].value = value
		indexedNums[index].index = index
	}

	sort.Slice(indexedNums, func(i, j int) bool {
		return indexedNums[i].value < indexedNums[j].value
	})

	i, j := 0, len(indexedNums)-1
	for i < j {
		if indexedNums[i].value+indexedNums[j].value == target {
			return []int{indexedNums[i].index, indexedNums[j].index}
		} else if indexedNums[i].value+indexedNums[j].value < target {
			i++
		} else {
			j--
		}
	}

	return []int{-1, -1}
}
