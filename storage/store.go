package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/huyvu/go-fiber-todolist/models"
)

var Users []models.User

const dataFile = "data/users.json"

func LoadData() {
	file, err := os.Open(dataFile)
	if err != nil {
		log.Println("Creating new data file")
		Users = []models.User{}
		return
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &Users)
}

func SaveData() {
	bytes, err := json.MarshalIndent(Users, "", "  ")
	if err != nil {
		log.Println("Failed to save data:", err)
		return
	}
	ioutil.WriteFile(dataFile, bytes, 0644)
}
