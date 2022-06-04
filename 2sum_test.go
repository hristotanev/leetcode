package main

import "testing"

func Test2SumExample1(t *testing.T) {
	res := TwoSum([]int{2, 7, 11, 15}, 9)
	if res[0] != 0 && res[1] != 1 {
		t.Fail()
	}
}

func Test2SumExample2(t *testing.T) {
	res := TwoSum([]int{3, 2, 4}, 6)
	if res[0] != 1 && res[1] != 2 {
		t.Fail()
	}
}

func Test2SumExample3(t *testing.T) {
	res := TwoSum([]int{3, 3}, 6)
	if res[0] != 0 && res[1] != 1 {
		t.Fail()
	}
}
