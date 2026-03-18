package steps

import (
	"errors"
	"fmt"
)

type DividerError struct {
	a, b int
}

func (d *DividerError) Error() string {
	return fmt.Sprintf("from Divider can`t divide by zero: %d / %d ", d.a, d.b)
}

func Divade(a, b int) (float64, error) {
	if b == 0 {
		return 0, &DividerError{a, b}
	} else {
		return float64(a) / float64(b), nil
	}
}

func MakeErrors() {

	fres, ferr := divide(10, 20)
	if ferr != nil {
		fmt.Println("error: ", ferr)
	} else {
		fmt.Println("result: ", fres)
	}

	sres, serr := divide(10, 0)
	if serr != nil {
		fmt.Println("error: ", serr)
	} else {
		fmt.Println("result: ", sres)
	}
	fmt.Println("......")
	res, err := Divade(20, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result: ", res)
	}
}

func divide(lhs, rhs int) (float64, error) {

	if rhs == 0 {
		return 0, errors.New("can`t divide by zero")
	} else {

		return float64(lhs) / float64(rhs), nil
	}
}
