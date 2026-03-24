package steps

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func round(a int, b float32) int {
	return a + int(math.Round(float64(b)))
}

func multyreturn() []interface{} {
	return []interface{}{1, 2, "test", "test2", nil, nil}
}

func getUserRank() (string, error) {
	randomNumber := rand.Intn(2)

	if randomNumber%2 == 0 {
		return "admin", nil
	}
	return "", errors.New("bad result")
}

func MakeFunctions() {
	fmt.Println("--------Functions-------\n--------Multy return--------")
	fmt.Println(round(5, 3.9))
	fmt.Println(multyreturn())
	el1 := multyreturn()

	fmt.Println(el1)

	for _, v := range el1 {

		fmt.Printf("%v ", v)
	}
	fmt.Println("")

	user, err := getUserRank()

	if err != nil {
		fmt.Println((err))
	}

	fmt.Println(user)

	fmt.Println("--------	-------")
}
