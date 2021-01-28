package main

import (
	"errors"
	"fmt"
	"sort"
)

func NextNumber(num int) int {
	seq := buildSeq(num)
	s, e := 0, len(seq)-1
	minIdx := e
	for s <= e {
		if seq[e] < seq[minIdx] {
			seq[e], seq[minIdx] = seq[minIdx], seq[e]
			break
		} else {
			e -= 1
		}
	}
	if e < 0 {
		return -1
	}
	sort.Ints(seq[e+1:])
	return buildInt(seq)
}
func buildInt(seq []int) int {
	res := 0
	for _, a := range seq {
		res *= 10
		res += a
	}
	return res
}
func buildSeq(num int) []int {
	if num == 0 {
		return []int{num}
	}
	seq := []int{}
	for num > 0 {
		seq = append([]int{num % 10}, seq...)
		num /= 10
	}
	return seq
}

func main() {
	tmp := 123456
	for i := 0; i < 50; i += 1 {
		fmt.Println(tmp)
		tmp = NextNumber(tmp)
	}
}
