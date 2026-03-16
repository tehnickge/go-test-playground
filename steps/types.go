package steps

import (
	"firstapp/lib"
	"firstapp/types"
	"fmt"
)

func MakeTypes() {
	fmt.Println("--------Types-------")
	user := types.User{
		Id:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Rank:  lib.StringPtr("admin"),
	}

	fmt.Println(user.PrettiePrint())
	fmt.Println(user.Greet())
	fmt.Println("user is admin:", user.IsAdmin())

	fmt.Println("---------------")

}
