package steps

import (
	"firstapp/lib"
	"fmt"
)

func MakeFunctions() {
	fmt.Println("--------Functions-------\n--------Multy return--------")
	fmt.Println(lib.Test(5, 3.9))
	fmt.Println(lib.Multyreturn())
	el1 := lib.Multyreturn()

	fmt.Println(el1)

	for _, v := range el1 {

		fmt.Printf("%v ", v)
	}
	fmt.Println("")

	user, err := lib.GetUserRank()

	if err != nil {
		fmt.Println((err))
	} else {
		fmt.Println(user)
	}

	fmt.Println("--------	-------")
}
