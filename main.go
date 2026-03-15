package main

import (
	"fmt"
	"math"
)

func main() {
	str := "GO🔥"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	printRunes(str)

	fmt.Println(test(5, 3.9))

	fmt.Println(multyreturn())
	el1 := multyreturn()

	fmt.Println(el1)

	for _, v := range el1 { 

		fmt.Printf("%v ", v)
	}

	

}

func printRunes(s string) {
	for _, r := range s {
		fmt.Printf("%c", r)
	}

	fmt.Println()
}

func test(a int, b float32) int {
	return a + int(math.Round(float64(b)))
}

func multyreturn() []interface{} {
	return []interface{}{1, 2, "sex", "sexxxxxxxxxxxxxxxx", nil, nil}
}
