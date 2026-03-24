package main

import (
	"bufio"
	"firstapp/internal/config"
	"firstapp/internal/models"
	"firstapp/internal/steps"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	config.InitDB()

	fmt.Println("Выполняется миграция таблиц...")

	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Printf("❌ Ошибка миграции: %v\n", err)
		return
	}

	fmt.Println("✅ Миграция завершена успешно!")
	fmt.Println("🚀 Приложение запущено!")

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
		fmt.Println("8. Interfaces")
		fmt.Println("9. Errors")
		fmt.Println("10. ReadAndWriteBuffer")
		fmt.Println("11. Generics")
		fmt.Println("12. Gorutine")
		fmt.Println("13. Mutex")
		fmt.Println("14. Context")
		fmt.Println("15. DB & ORM")
		fmt.Println("0. Exit")
		fmt.Println("Enter the number of the step you want to see: ")

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
		case 8:
			steps.MakeInterfaces()
		case 9:
			steps.MakeErrors()
		case 10:
			steps.MakeReadAndWrite()
		case 11:
			steps.MakeGanerics()
		case 12:
			steps.MakeGoRutine()
		case 13:
			steps.MakeMutex()
		case 14:
			steps.MakeContext()
		case 15:
			steps.MakeDB()
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid step number")
		}
		fmt.Println()
	}
}
