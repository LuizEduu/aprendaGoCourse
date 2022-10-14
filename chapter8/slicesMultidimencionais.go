package main

import "fmt"

func main() {

	sliceOfSlices := [][]int{
		// 0  1  2
		{1, 2, 3}, // 0

		//   0  1  2
		{4, 5, 6}, // 1
	}

	fmt.Println(sliceOfSlices[1][1])
}
