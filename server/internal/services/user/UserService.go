package UserService

import (
	"errors"
	UserModel "server/server/internal/models"
	Database "server/server/internal/services/database"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(user *UserModel.User) error {
	db := Database.GetInstance()

	//convert the password to a hashed password before insertion
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	//insert the user into the database
	res, err := db.Query("INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id;", user.Email, user.Username, user.Password)
	if err != nil {
		msg := strings.TrimPrefix(err.Error(), "pq:")
		msg = strings.TrimSpace(msg)

		return errors.New(msg)
	}
	defer res.Close()

	if res.Next() {
		err = res.Scan(&user.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func VerifyUser(creds *UserModel.Credentials) (int, error) {
	var id int = -1
	var hashedPassword string = ""

	db := Database.GetInstance()

	res, err := db.Query("SELECT id, password FROM users WHERE username = $1", creds.Username)
	if err != nil {
		return -1, err
	}
	defer res.Close()

	if res.Next() {
		err = res.Scan(&id, &hashedPassword)
		if err != nil {
			return -1, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
		if err != nil {
			return -1, errors.New("passwords did not match")
		}

		return id, nil
	} else {
		return -1, errors.New("user not found")
	}
}
