package steps

import (
	"firstapp/lib"
	"fmt"
)

func MakeStrings() {
	fmt.Println("--------Strings-------")
	str := "GO🔥"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	lib.PrintRunes(str)

	fmt.Println("--------	-------")
}
