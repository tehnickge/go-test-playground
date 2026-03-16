package steps

func changeValuePtr(val *int) {
	*val = 225
}

func MakePtrs() {
	var count int = 10
	var countPtr *int = &count

	println("Value:", count)
	println("Pointer:", countPtr)
	println("Dereferenced Pointer:", *countPtr)

	// Изменение значения через указатель
	*countPtr = 20

	println("Value:", count)
	println("Pointer:", countPtr)
	println("Dereferenced Pointer:", *countPtr)

	// Изменение значения напрямую
	count = 30

	println("Value:", count)
	println("Pointer:", countPtr)
	println("Dereferenced Pointer:", *countPtr)

	// Изменение значения через функцию, которая принимает указатель
	changeValuePtr(countPtr)

	println("Value:", count)
	println("Pointer:", countPtr)
	println("Dereferenced Pointer:", *countPtr)

	// Также можно передать указатель на переменную напрямую в функцию, которая ожидает указатель, например:
	changeValuePtr(&count)

	println("Value:", count)
	println("Pointer:", countPtr)
	println("Dereferenced Pointer:", *countPtr)

	// в этих примерах мы видим, что изменение значения через указатель изменяет исходную переменную,
	// и наоборот, изменение переменной напрямую также отражается при доступе через указатель.
	// Это демонстрирует взаимосвязь между переменной и её указателем в Go.
	// & - оператор получения адреса(как ссылка в c++),
	// * - оператор разыменования указателя и оператор изменения значения по указателю(как разыменование в c++).
	// * - также используется для объявления указателя на определённый тип данных, например, *int - указатель на целое число.

}
