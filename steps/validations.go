package steps

import (
	"fmt"
	"regexp"
)

func MakeValidations() {
	var email string
	fmt.Print("Enter an email address: ")
	fmt.Scanln(&email)
	fmt.Printf("Email is valid: %t\n", emailValidation(email))
}

func emailValidation(email string) bool {
	res, err := regexp.Compile(`.+@.+\..+`)
	if err != nil {
		panic(fmt.Sprintf("failed to compile regexp: %v", err))
	}
	return res.MatchString(email)
}
