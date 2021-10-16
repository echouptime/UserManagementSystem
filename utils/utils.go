package utils

import (
	"UserManagementSystem/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Input(prompt string) string {
	var text string
	fmt.Printf(prompt)
	fmt.Scan(&text)
	return text
}

func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
		return false
	}
}

func SaveDb(tasks []*models.User) {
	os.Remove("userInfo.json")
	file, err := os.OpenFile("userInfo.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	jsonCode := json.NewEncoder(file)
	jsonCode.Encode(&tasks)
}
