package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	// // Seeding with the same value results in the same random sequence each run.
	// // For different numbers, seed with a different value, such as
	// // time.Now().UnixNano(), which yields a constantly-changing number.
	// rand.Seed(86)
	// fmt.Println(rand.Intn(100))
	// fmt.Println(rand.Intn(100))
	// fmt.Println(rand.Intn(100))

	//
	//

	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)

	s1 := arr[1:4]
	fmt.Println(s1)

	for i, valor := range arr {
		fmt.Println(i, valor)
	}

	s2 := make([]int, 10, 20)
	fmt.Println(s2, len(s2), cap(s2))

	s3 := s2[:10]
	fmt.Println(s3, len(s3), cap(s3))

	s4 := make([]int, 10, 10)
	copy(s4, s3)
	fmt.Println(s4, len(s4), cap(s4))
}
