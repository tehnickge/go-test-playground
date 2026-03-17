package lib

import "strconv"

func StringPtr(s string) *string {
	return &s
}

func StringToInt(s string) (int, error) {
	// Здесь можно использовать strconv.Atoi для конвертации строки в целое число
	// Например:
	result, err := strconv.Atoi(s) // Это вернет (int, error), нужно обработать результат и вернуть его
	return result, err
}
