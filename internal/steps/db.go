package steps

import (
	"firstapp/internal/config"
	"firstapp/internal/models"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func MakeDB() {
	var name string
	var userIdToDelete uint

	fmt.Print("Enter user name: ")
	fmt.Scanln(&name)

	user, err := addUser(name)
	if err != nil {
		fmt.Println("❌ Error creating user:", err)
		return
	}

	fmt.Println("✅ Created user:", user.ID, user.Name, user.UUID, user.Email)

	// дальше работа
	users, err := getUsers()
	if err != nil {
		fmt.Println("❌ Error getting users:", err)
		return
	}

	fmt.Println("📦 All users:")
	for _, u := range users {
		fmt.Println("📦 All users:", u.ID, u.Name, u.UUID, u.Email)
	}

	userFromBD, error := getUserByID(user.ID)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("user from get UserById: ", userFromBD.ID, userFromBD.Name, userFromBD.UUID, userFromBD.Email)

	fmt.Print("Eneter user id to delete: ")
	fmt.Scanln(&userIdToDelete)
	deleteUser(userIdToDelete)

}

func addUser(name string) (models.User, error) {
	user := models.User{
		Name:     name,
		Email:    strings.ToLower(fmt.Sprintf("%s@example.com", name)),
		UUID:     uuid.New().String(),
		Password: "124343",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func getUsers() ([]models.User, error) {
	var users []models.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		fmt.Println("❌ Error getting users:", result.Error)
		return []models.User{}, result.Error
	}

	return users, nil
}

func getUserByID(id uint) (models.User, error) {
	var user models.User

	result := config.DB.First(&user, id)

	if result.Error != nil {
		fmt.Println("❌ Error getting user:", result.Error)
		return models.User{}, result.Error
	}

	return user, nil
}

func deleteUser(id uint) {
	result := config.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		fmt.Println("❌ Error deleting user:", result.Error)
		return
	}

	fmt.Println("🗑 User deleted:", id)
}
