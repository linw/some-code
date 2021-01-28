package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	s, e := 0, length(arr)-1
	for s <= e {
		m := (s + e) / 2
		if arr[m] == target {
			return m
		}
		if arr[m] > target {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return -1
}
func SpecialSearch(arr []int, target int) int {
	s := 0
	e := len(arr) - 1
	for s <= e {
		m := (s + e) / 2
		if target == arr[m] {
			return m
		} else {
			if arr[s] <= arr[m] {
				if target >= arr[s] && target < arr[m] {
					e = m - 1
				} else {
					s = m + 1
				}
			} else {
				if target > arr[m] && target <= arr[e] {
					s = m + 1
				} else {
					e = m - 1
				}
			}
		}
	}
	return -1
}

func main() {
	arr := []int{4, 5, 6, 9, 10, 1, 3}
	targets := []int{5, 10, 3, 9, -1}
	for _, target := range targets {
		fmt.Println(arr, target, SpecialSearch(arr, target))
	}
}
