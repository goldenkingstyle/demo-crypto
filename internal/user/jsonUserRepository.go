package user

import (
	"encoding/json"
	"log"
	"os"
)

type JsonUserRepository struct {
	filepath string
}

// TODO: error handle
func (r *JsonUserRepository) Get() *User {
	userJson, err := os.ReadFile(r.filepath)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}

func (r *JsonUserRepository) Save(user *User) {
	userJson, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(r.filepath, userJson, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func NewJsonUserRepository(filepath string) *JsonUserRepository {
	return &JsonUserRepository{filepath}
}
