package main

import "fmt"

func main() {
	arr := []int{2, 5, 9, 15, 20, 25, 37, 45, 50, 55, 71}
	fmt.Println(binary_Search(arr, 37))

}

func binary_Search(arr []int, x int) int {
	L := 0
	R := len(arr) - 1
	for L <= R {
		Mid := (L + R) / 2
		if arr[Mid] == x {
			L = Mid + 1
		} else if arr[Mid] > x {
			R = Mid - 1
		} else {
			return Mid
		}
	}
	return -1
}
