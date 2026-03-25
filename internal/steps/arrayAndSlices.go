package steps

import (
	"firstapp/internal/constants"
	"fmt"
)

func MakeArrayAndSlices() {
	fmt.Println("--------Array and Slices-------")
	fmt.Println(constants.ArrayAny)

	for _, v := range constants.ArrayAny {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	var testArray [5]int = [5]int{1, 2, 3, 4, 5}
	var testArray1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 55}
	var testArray2 []int = []int{229, 5543, 5435, 2342}
	testArray1 = append(testArray1, testArray[:]...)
	testArray1 = append(testArray1, testArray2...)
	testArray2 = append(testArray2, append(testArray1, testArray[:]...)...)
	fmt.Println(testArray[:])
	fmt.Println(testArray1[:])
	fmt.Println(testArray2[:])
	fmt.Println("---------------")
}
