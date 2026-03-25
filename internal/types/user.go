package types

import (
	"encoding/json"
	"fmt"
)

// Структура для представления пользователя
type User struct {
	Id    int
	Name  string
	Email string
	Rank  *string
}

// Метод для структуры User
func (u *User) Greet() string {
	return fmt.Sprintf("Hello, my name is %s!", u.Name)
}

func (u *User) IsAdmin() bool {
	return u.Rank != nil && *u.Rank == "admin"
}

func (u *User) PrettiePrint() string {
	jsonData, err := json.MarshalIndent(u, "", "	")
	if err != nil {
		return "Error occurred while marshaling user data"
	}
	return string(jsonData)
}

func (u *User) ChangeName(newName string) {
	u.Name = newName
}

func (u User) CopyUser(otherUser User) User {
	u.Name = otherUser.Name

	if otherUser.Rank != nil {
		newRank := *otherUser.Rank
		u.Rank = &newRank
	} else {
		u.Rank = nil
	}

	return u
}

func (u User) Clone() User {
	clone := u

	if u.Rank != nil {
		newRank := *u.Rank
		clone.Rank = &newRank
	} else {
		clone.Rank = nil
	}

	return clone
}

var Userino = struct {
	Id   int
	Name string
}{
	Id:   1,
	Name: "John",
}
