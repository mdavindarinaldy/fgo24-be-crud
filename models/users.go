package models

import (
	"backend2/utils"
	"context"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Name     string `form:"name" json:"name" db:"name" binding:"required"`
	Email    string `form:"email" json:"email" db:"email" binding:"required,email"`
	Password string `form:"password" json:"password" db:"password"`
}

type ResponseUser struct {
	Name  string `form:"name" json:"name" db:"name" binding:"required"`
	Email string `form:"email" json:"email" db:"email" binding:"required,email"`
}

func FindAllUsers() ([]ResponseUser, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []ResponseUser{}, err
	}
	defer conn.Close()
	rows, err := conn.Query(context.Background(), `SELECT name, email, password FROM users`)
	if err != nil {
		return []ResponseUser{}, err
	}
	users, err := pgx.CollectRows[ResponseUser](rows, pgx.RowToStructByName)
	if err != nil {
		return []ResponseUser{}, err
	}
	return users, nil
}

func FindUser(idParam string) ([]User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}
	defer conn.Close()
	id, _ := strconv.Atoi(idParam)
	rows, err := conn.Query(context.Background(), `SELECT name, email FROM users WHERE id = $1`, id)
	if err != nil {
		return []User{}, err
	}
	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func CreateUser(user User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(context.Background(), `INSERT INTO users (name, email, password) VALUES ($1,$2,$3)`, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(idParam string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	id, _ := strconv.Atoi(idParam)
	_, err = conn.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
