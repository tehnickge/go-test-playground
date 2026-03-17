package steps

import "fmt"

type Item interface {
	WhoIAm() string
}

type Gun int

func (m Gun) WhoIAm() string {
	return fmt.Sprintf("gun: %d", m)
}

type AnotherType struct{}

func (a AnotherType) WhoIAm() string {
	return "AnotherType"
}

func execute(i Item) {
	println(i.WhoIAm())
}

func MakeInterfaces() {
	var value Gun = Gun(10)
	var anotherValue AnotherType = AnotherType{}

	execute(value)
	execute(anotherValue)
}
