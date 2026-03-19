package steps

import (
	"cmp"
	"fmt"
	"slices"
)

type TestItem struct {
	name string
}

// ~ - означает, что любой тип, который является, допустим, int32 будет автомотически приведен к нему
// пример type Speed int32 - если мы поместим в функцию с ограниченим, где будет указано только Int32 - она отработает \
// логично, что без знака "~" компилятор выдаст ошибку
type Number32 interface {
	~int32 | ~uint32
}
type Speed int32

type MyArray[T cmp.Ordered] struct {
	inner []T
}

func MakeGanerics() {

	fmt.Println(IsEqual(1, 2))
	fmt.Println(IsEqual(1, 1))
	fmt.Println(IsEqual(TestItem{name: "testName"}, TestItem{name: "testName"}))

	arr := []Speed{Speed(1), Speed(23), Speed(-400)}
	fmt.Println(sumNumbers(arr))
	arr1 := []uint32{uint32(1), uint32(23), uint32(400)}
	fmt.Println(sumNumbers(arr1))

	arr2 := MyArray[int32]{inner: []int32{132422, 2321, 3321, -32, 23, 1, 0}}
	fmt.Println(arr2)
	fmt.Println(arr2.Max())
	arr2.Sort()
	fmt.Println(arr2)
}

func IsEqual[T comparable](a, b T) bool {

	return a == b
}

func sumNumbers[T Number32](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}

	return sum
}

func (m *MyArray[T]) MaxBasic() T {
	if len(m.inner) == 0 {
		var zero T
		return zero
	}
	max := m.inner[0]
	for _, v := range m.inner {
		if v > max {
			max = v
		}
	}
	return max
}

func (m *MyArray[T]) Max() T {
	if len(m.inner) == 0 {
		var zero T
		return zero
	}
	return slices.Max(m.inner)
}

func (m *MyArray[T]) Sort() {
	slices.Sort(m.inner)
}
