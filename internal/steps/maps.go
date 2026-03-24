package steps

import (
	"firstapp/constants"
	"fmt"
)

func MakeMaps() {
	fmt.Println(constants.TestMap)
	// удаление элемента по ключу
	delete(constants.TestMap, "token")
	fmt.Println(constants.TestMap)

	// поиск элемента по ключу
	fmt.Print("Enter a key to search: ")
	// Считываем строку с клавиатуры
	var searchedString string
	fmt.Scanln(&searchedString)
	res, found := constants.TestMap[searchedString]
	if !found {
		fmt.Println("key not found")
	} else {
		fmt.Println(res, found)
	}
	fmt.Println("---------------")

	// перебор элементов в map
	for key, value := range constants.TestMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
