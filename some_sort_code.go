package main

import "fmt"

func QuickSort(arr []int, s int, e int) {
	if s >= e {
		return
	}
	base, i, j := arr[s], s, e
	for i < j {
		for arr[j] >= base && i < j {
			j -= 1
		}
		for arr[i] <= base && i < j {
			i += 1
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[s], arr[j] = arr[j], arr[s]
	QuickSort(arr, s, j-1)
	QuickSort(arr, j+1, e)
}

func HeapSort(arr []int, s int, e int) {

}

func main() {
	arr := []int{-10, 9, 28, 5, 2, 4, 10}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	HeapSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
