package user

import (
	"database/sql"
	"fmt"

	"github.com/Tibz-Dankan/reserve-now-microservices/internal/config"
)

type User struct {
	UserId      int    `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var db = config.Db()

func FindOne(userId int) (User, error) {
	var user User
	err := db.QueryRow(`SELECT "userId", "firstName", "lastName","phoneNumber",
               "createdAt","updatedAt" FROM _users WHERE "userId" = $1`,
		userId).Scan(&user.UserId, &user.FirstName, &user.LastName,
		&user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user of provided id is not found")
		}
		return user, fmt.Errorf("error occurred while querying table")
	}

	return user, nil
}

func FindAll() ([]User, error) {
	rows, err := db.Query(`SELECT "userId", "firstName", "lastName","phoneNumber",
                 "createdAt","updatedAt" FROM _users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var usr User
		err := rows.Scan(&usr.UserId, &usr.FirstName, &usr.LastName,
			&usr.PhoneNumber, &usr.CreatedAt, &usr.UpdatedAt)

		if err != nil {
			return users, err
		}
		users = append(users, usr)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
