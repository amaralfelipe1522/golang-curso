package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	// append e copy são apenas para slices
	slice1 = append(slice1, 4, 5, 6)

	fmt.Println(array1, slice1)

	slice2 := make([]int, 2, 10)
	copy(slice2, slice1)
	fmt.Println(slice2, cap(slice2))

}
