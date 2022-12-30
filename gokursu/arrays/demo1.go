package arrays

import "fmt"

/*
func Demo2() {

	nums := [4]int{1, 2, 3, 4}
	Max := nums[0]
	for i := 0; i < len(nums); i++ {

		if Max < nums[i] {
			Max = nums[i]

		}

	}
	fmt.Println("Max number is ", Max)
}

*/
func Demo2() {

	nums := [10]int{50, 73, 92, 846, 41, 658, 741, 654, 98, 746}
	Max := nums[0]

	for i := 0; i < len(nums); i++ {

		if Max < nums[i] {
			Max = nums[i]
		}
	}
	fmt.Println("max number is  ", Max)
}
