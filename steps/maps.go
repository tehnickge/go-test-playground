package steps

import (
	"bufio"
	"firstapp/constants"
	"fmt"
	"os"
	"strings"
)

func MakeMaps() {
	fmt.Println(constants.TestMap)
	delete(constants.TestMap, "token")
	fmt.Println(constants.TestMap)
	fmt.Println("---------------")
	fmt.Print("Enter a key to search: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchedString := strings.TrimSpace(scanner.Text())
	res, found := constants.TestMap[searchedString]
	if !found {
		fmt.Println("key not found")
	} else {
		fmt.Println(res, found)
	}
	fmt.Println("---------------")

	for key, value := range constants.TestMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
