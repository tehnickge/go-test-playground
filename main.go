package main

import (
	"firstapp/constants"
	"firstapp/steps"
	"fmt"
)

func main() {
	steps.MakeStrings()
	steps.MakeFunctions()
	steps.MakeTyepes()

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

	fmt.Println(constants.TestMap)
	delete(constants.TestMap, "token")
	fmt.Println(constants.TestMap)
	fmt.Println("---------------")
	fmt.Print("Enter a key to search: ")
	var searchedString string
	fmt.Scanf("%s", &searchedString)
	res, found := constants.TestMap[searchedString]
	if !found {
		fmt.Println("key not found")
	}
	fmt.Println(res, found)
	fmt.Println("---------------")

	for key, value := range constants.TestMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
