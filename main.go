package main

import (
	"bufio"
	"firstapp/steps"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Selected steps:")
		fmt.Println("1. Strings")
		fmt.Println("2. Functions")
		fmt.Println("3. Types")
		fmt.Println("4. Arrays and Slices")
		fmt.Println("5. Maps")
		fmt.Println("6. Pointers")
		fmt.Println("7. Validations")
		fmt.Println("0. Exit")
		fmt.Print("Enter the number of the step you want to see: ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		stepNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue
		}

		switch stepNumber {
		case 1:
			steps.MakeStrings()
		case 2:
			steps.MakeFunctions()
		case 3:
			steps.MakeTypes()
		case 4:
			steps.MakeArrayAndSlices()
		case 5:
			steps.MakeMaps()
		case 6:
			steps.MakePtrs()
		case 7:
			steps.MakeValidations()
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid step number")
		}
		fmt.Println()
	}
}
