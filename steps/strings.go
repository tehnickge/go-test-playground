package steps

import (
	"firstapp/lib"
	"fmt"
)

func MakeStrings() {
	fmt.Println("--------Strings-------")

	str := "GO🔥"

	// перебор строки байтам
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	// перебор строки рунам
	lib.PrintRunes(str)

	fmt.Println("--------	-------")
}
