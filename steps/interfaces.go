package steps

import (
	"fmt"
)

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

type Player interface {
	CanPlay() bool
	GetName() string
	SetName(name string)
	SetNameWithOutPtr(name string)
	CloneWithNewName(name string) Player
}

type Dog struct {
	name string
}

func (d *Dog) CanPlay() bool {
	return true
}
func (d *Dog) GetName() string {
	return d.name
}
func (d *Dog) SetName(name string) {
	d.name = name
}

// запись не будет происходить т.к. мы работаем с копией
func (d Dog) SetNameWithOutPtr(name string) {
	d.name = name
}
func (d Dog) CloneWithNewName(name string) Player {
	newDog := d
	newDog.name = name
	return &newDog
}

type Cat struct {
	name string
}

func (d *Cat) CanPlay() bool {
	return true
}
func (d *Cat) GetName() string {
	return d.name
}
func (d *Cat) SetName(name string) {
	d.name = name
}

// запись не будет происходить т.к. мы работаем с копией
func (d Cat) SetNameWithOutPtr(name string) {
	d.name = name
}
func (d Cat) CloneWithNewName(name string) Player {
	newDog := d
	newDog.name = name
	return &newDog
}

func MakeInterfaces() {
	var value Gun = Gun(10)
	var anotherValue AnotherType = AnotherType{}

	// В Go интерфейсы являются ссылочными типами,
	// и когда вы передаете значение типа, который реализует интерфейс,
	// вы фактически передаете копию этого значения.
	// Это означает, что если вы измените значение внутри функции,
	// оно не повлияет на оригинальное значение за пределами функции.
	execute(value)
	execute(anotherValue)
	// Если вы хотите передать указатель на значение,
	// то изменения внутри функции будут отражаться на оригинальном значении.
	// Это связано с тем, что указатель позволяет функции работать с оригинальным
	// значением, а не с его копией.
	execute(&value)
	execute(&anotherValue)

	var dog Player = &Dog{name: "Buddy"}
	var cat Player = &Cat{name: "PDiddy"}
	fmt.Println(dog)
	fmt.Println(dog.GetName())
	fmt.Printf("dog can play? %t\n", dog.CanPlay())
	dog.SetName("luna")
	fmt.Println(dog.GetName())
	dog.SetNameWithOutPtr("nikita")
	fmt.Println(dog.GetName())
	var clone Player = dog.CloneWithNewName("tesla")
	fmt.Println(clone.GetName())
	fmt.Println(dog.GetName())
	fmt.Println("check struct with downcast")
	fmt.Printf("cat name: %v\ndog name: %v\n", cat.GetName(), dog.GetName())
	changeNameForDog(cat)
	changeNameForDog(dog)
	fmt.Printf("cat name: %v\ndog name: %v\n", cat.GetName(), dog.GetName())
	fmt.Println("check struct with downcast")
	changeNameForAnotherPlaye(dog)
	changeNameForAnotherPlaye(cat)
	fmt.Printf("cat name: %v\ndog name: %v\n", cat.GetName(), dog.GetName())

}

func changeNameForDog(p Player) {
	if player, ok := p.(*Dog); ok {
		player.SetName("testName")
	} else {
		fmt.Printf("this player not a dog %T\n", p)
	}
}

func changeNameForAnotherPlaye(p Player) {
	switch v := p.(type) {
	case *Dog:
		v.SetName("dog")
	case *Cat:
		v.SetName("cat")
	}

}
