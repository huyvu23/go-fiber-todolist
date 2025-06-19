package services

import (
	"github.com/huyvu/go-fiber-todolist/models"
	"github.com/huyvu/go-fiber-todolist/storage"
)

func GetAllUsers() []models.User {
	return storage.Users
}

func CreateUser(user models.User) models.User {
	user.ID = getNextID()
	storage.Users = append(storage.Users, user)
	storage.SaveData()
	return user
}

func getNextID() int {
	maxID := 0
	for _, u := range storage.Users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	return maxID + 1
}
