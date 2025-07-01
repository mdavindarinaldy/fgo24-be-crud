package models

import (
	"backend2/utils"
	"context"
	"encoding/json"
	"fmt"

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
	result := utils.RedistClient.Exists(context.Background(), "/users")
	if result.Val() == 0 {
		conn, err := utils.DBConnect()
		if err != nil {
			return []ResponseUser{}, err
		}
		defer conn.Close()
		rows, err := conn.Query(context.Background(), `SELECT name, email FROM users`)
		if err != nil {
			return []ResponseUser{}, err
		}
		users, err := pgx.CollectRows[ResponseUser](rows, pgx.RowToStructByName)
		if err != nil {
			return []ResponseUser{}, err
		}

		encoded, err := json.Marshal(users)
		if err != nil {
			return []ResponseUser{}, err
		}

		utils.RedistClient.Set(context.Background(), "/users", string(encoded), 0)
		return users, nil
	} else {
		data := utils.RedistClient.Get(context.Background(), "/users")
		str := data.Val()
		users := []ResponseUser{}
		json.Unmarshal([]byte(str), &users)
		return users, nil
	}
}

func FindUser(id int) (ResponseUser, error) {
	redisEndpoint := fmt.Sprintf("/users:%d", id)
	result := utils.RedistClient.Exists(context.Background(), redisEndpoint)
	if result.Val() == 0 {
		conn, err := utils.DBConnect()
		if err != nil {
			return ResponseUser{}, err
		}
		defer conn.Close()
		rows, err := conn.Query(context.Background(), `SELECT name, email FROM users WHERE id = $1`, id)
		if err != nil {
			return ResponseUser{}, err
		}
		user, err := pgx.CollectOneRow[ResponseUser](rows, pgx.RowToStructByName)
		if err != nil {
			return ResponseUser{}, err
		}
		encoded, err := json.Marshal(user)
		if err != nil {
			return ResponseUser{}, err
		}
		utils.RedistClient.Set(context.Background(), redisEndpoint, string(encoded), 0)
		return user, nil
	} else {
		data := utils.RedistClient.Get(context.Background(), redisEndpoint)
		str := data.Val()
		user := ResponseUser{}
		json.Unmarshal([]byte(str), &user)
		return user, nil
	}
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
	result := utils.RedistClient.Exists(context.Background(), "/users")
	if result.Val() != 0 {
		utils.RedistClient.Del(context.Background(), "/users")
	}
	return nil
}

func UpdateUser(id int, user User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(
		context.Background(),
		`UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`,
		user.Name, user.Email, user.Password, id,
	)
	if err != nil {
		return err
	}

	redisEndpoint := fmt.Sprintf("/users:%d", id)
	result := utils.RedistClient.Exists(context.Background(), redisEndpoint)
	if result.Val() != 0 {
		utils.RedistClient.Del(context.Background(), redisEndpoint)
	}

	return nil
}

func DeleteUser(id int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	redisEndpoint := fmt.Sprintf("/users:%d", id)
	result := utils.RedistClient.Exists(context.Background(), redisEndpoint)
	if result.Val() != 0 {
		utils.RedistClient.Del(context.Background(), redisEndpoint)
	}

	result = utils.RedistClient.Exists(context.Background(), "/users")
	if result.Val() != 0 {
		utils.RedistClient.Del(context.Background(), "/users")
	}

	return nil
}
