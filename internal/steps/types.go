package steps

import (
	"firstapp/internal/types"
	"firstapp/lib"
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

	// изменение имени пользователя
	user.ChangeName("Nikola Tesla")
	fmt.Println(user.PrettiePrint())

	// копирование структуры
	copiedUser := user.CopyUser(user)

	fmt.Printf(
		"Original user: %s | struct addr: %p | Rank pointer: %p\n",
		user.Name,
		&user,
		user.Rank,
	)

	fmt.Printf(
		"Copied user: %s | struct addr: %p | Rank pointer: %p\n",
		copiedUser.Name,
		&copiedUser,
		copiedUser.Rank,
	)

	cloneUser := user.Clone()

	fmt.Printf(
		"Clone user: %s | struct addr: %p | Rank pointer: %p\n",
		cloneUser.Name,
		&cloneUser,
		cloneUser.Rank,
	)

	fmt.Println("---------------")

}
