package main

import (
	"encoding/json"
	"firstapp/constants"
	"firstapp/lib"
	"firstapp/types"
	"fmt"
)

func main() {
	str := "GO🔥"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	lib.PrintRunes(str)

	fmt.Println(lib.Test(5, 3.9))

	fmt.Println(lib.Multyreturn())
	el1 := lib.Multyreturn()

	fmt.Println(el1)

	for _, v := range el1 {

		fmt.Printf("%v ", v)
	}
	fmt.Println()
	user, err := lib.GetUserRank()

	if err != nil {
		fmt.Println((err))
	} else {
		fmt.Println(user)
	}

	data := types.User{
		Id:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Rank:  lib.StringPtr("admin"),
	}

	jsonData, _ := json.MarshalIndent(data, "", "	")
	fmt.Println(string(jsonData))

	userino, errino := json.MarshalIndent(types.Userino, "", "	")

	fmt.Println(string(userino), errino)

	fmt.Println("---------------")
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
