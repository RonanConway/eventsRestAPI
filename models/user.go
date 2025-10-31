package models

import (
	"errors"

	"github.com/RonanConway/eventsRestAPI/db"
	"github.com/RonanConway/eventsRestAPI/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

var users = []User{}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// We don't want the password stored in plain text. Going to use a
	// hash function so that it scrambles it so it can't be recovered.
	// Then when a user logs in I will compare against the hash and not
	// the actual password.
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	// Becuase ID is AUTOINCREMENT we I need to update the userId with
	// the value we got back from the database
	user.ID = userId

	return err
}

// Compare the password passed in by the login to the hashed
// password stored in the DB and see if they match. Must
// has the password passed in the login and compare against the hashed
// DB password in order to make the comparison.
func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(retrievedPassword, user.Password)
	if !passwordIsValid {
		return errors.New("Invalid credentials: Password mistmatch!")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
