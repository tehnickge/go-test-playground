package types

type User struct {
	Id    int
	Name  string
	Email string
	Rank  *string
}

var Userino = struct {
	Id   int
	Name string
}{
	Id:   1,
	Name: "John",
}
