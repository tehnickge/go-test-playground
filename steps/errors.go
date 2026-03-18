package steps

import (
	"errors"
	"fmt"
)

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

}

func divide(lhs, rhs int) (float64, error) {

	if rhs == 0 {
		return 0, errors.New("can`t divide by zero")
	} else {

		return float64(rhs / lhs), nil
	}
}
