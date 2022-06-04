package main

import "testing"

func Test2SumExample1(t *testing.T) {
	res := TwoSum([]int{2, 7, 11, 15}, 9)

	expectedRes := []int{0, 1}
	if res[0] != expectedRes[0] || res[1] != expectedRes[1] {
		t.Errorf("Expected %x, but got %x.", expectedRes, res)
	}
}

func Test2SumExample2(t *testing.T) {
	res := TwoSum([]int{3, 2, 4}, 6)

	expectedRes := []int{1, 2}
	if res[0] != expectedRes[0] || res[1] != expectedRes[1] {
		t.Errorf("Expected %x, but got %x.", expectedRes, res)
	}
}

func Test2SumExample3(t *testing.T) {
	res := TwoSum([]int{3, 3}, 6)

	expectedRes := []int{0, 1}
	if res[0] != expectedRes[0] || res[1] != expectedRes[1] {
		t.Errorf("Expected %x, but got %x.", expectedRes, res)
	}
}

func Test2SumExample4(t *testing.T) {
	res := TwoSum([]int{3, 3}, 7)

	expectedRes := []int{-1, -1}
	if res[0] != expectedRes[0] || res[1] != expectedRes[1] {
		t.Errorf("Expected %x, but got %x.", expectedRes, res)
	}
}
