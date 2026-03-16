package lib

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func Add(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) int {
	return a * b
}

func PrintRunes(s string) {
	for _, r := range s {
		fmt.Printf("%c", r)
	}

	fmt.Println()
}

func Test(a int, b float32) int {
	return a + int(math.Round(float64(b)))
}

func Multyreturn() []interface{} {
	return []interface{}{1, 2, "sex", "sexxxxxxxxxxxxxxxx", nil, nil}
}

func GetUserRank() (string, error) {
	randomNumber := rand.Intn(2)

	if randomNumber%2 == 0 {
		return "admin", nil
	}
	return "", errors.New("bad result")
}

func GetTestResultOrError() (string, error) {

	randomNumber := rand.Intn(2)

	if randomNumber%2 == 0 {
		return "ok", nil
	}

	return "", errors.New("bad result")
}

func StringPtr(s string) *string {
	return &s
}