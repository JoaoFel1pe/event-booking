package models

import "event-booking-api/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES ($1, $2) RETURNING id`

	err := db.DB.QueryRow(query, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}
