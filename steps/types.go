package steps

import (
	"firstapp/lib"
	"firstapp/types"
	"fmt"
)

func MakeTypes() {
	fmt.Println("--------Types-------")
	// создание экземпляра структуры
	user := types.User{
		Id:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Rank:  lib.StringPtr("admin"),
	}

	// использование методов структуры
	fmt.Println(user.PrettiePrint())

	// вызов метода Greet
	fmt.Println(user.Greet())

	// проверка на админский ранг
	fmt.Println("user is admin:", user.IsAdmin())

	fmt.Println("---------------")

}
