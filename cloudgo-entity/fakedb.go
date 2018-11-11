package cloudgo_entity

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

//A fake Database that uses JSON to store user data



type UserInfo struct {
	Username string `json:"username"`
	StudentID string `json:"sid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type User struct {
	Username string `json:"username"`
	StudentID string `json:"sid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}


func GetUsers() []User {
	userList, err := os.Open("./users.json")
	if err != nil {
		return []User{}
	}
	decoder := json.NewDecoder(userList)
	res := make([]User, 0, 0)
	err = decoder.Decode(&res)
	if err != nil {
		return []User{}
	}
	return res
}

func addUser(newUser User) error {
	currentUsers := GetUsers()
	currentUsers = append(currentUsers, newUser)
	userList, err := os.Create("./users.json")
	if err != nil {
		return errors.New("fails to write json file")
	}
	encoder := json.NewEncoder(userList)
	err = encoder.Encode(currentUsers)
	return err
}

func CheckDuplicate(newUser User) bool {
	currentUsers := GetUsers()
	for _, user := range currentUsers {
		if user.Username == newUser.Username || user.Email == newUser.Email ||
			user.Phone == newUser.Phone || user.StudentID == newUser.StudentID {
				return false
		}
	}
	return true
}

func CheckPhoneDuplicate(phone string) bool {
	currentUsers := GetUsers()
	for _, user := range currentUsers {
		if user.Phone == phone {
			return false
		}
	}
	return true
}

func CheckEmailDuplicate(email string) bool {
	currentUsers := GetUsers()
	for _, user := range currentUsers {
		if user.Email == email {
			return false
		}
	}
	return true
}

func CheckUsernameDuplicate(username string) bool {
	currentUsers := GetUsers()
	for _, user := range currentUsers {
		if user.Username == username {
			return false
		}
	}
	return true
}

func CheckIDDuplicate(id string) bool {
	currentUsers := GetUsers()
	for _, user := range currentUsers {
		if user.StudentID == id {
			return false
		}
	}
	return true
}

func CheckSignin(name, pass string) int{
	currentUsers := GetUsers()
	for _, user := range currentUsers {
		if user.Username == name {
			if user.Password == pass {
				return 0
			} else {
				return 1
			}
		}
	}
	return 2
}

func GetUser(username string) *UserInfo{
	currentUsers := GetUsers()
	for _, user := range currentUsers {
		if user.Username == username {
			return &UserInfo{
				user.Username,
				user.StudentID,
				user.Email,
				user.Phone,
			}
		}
	}
	return nil
}

