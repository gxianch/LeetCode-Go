package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
		fmt.Println(m)
	}
	return nil
}
func twoSumG(nums []int,target int) []int {
	m :=make(map[int]int)
	for k , v :=range nums{
		if idx,ok :=m[target -v];ok{
			return []int{idx,k}
		}
		m[v] = k
	}
	return nil
}
