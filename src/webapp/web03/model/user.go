package model

import (
	"fmt"
	"webapp/web03/utils"
)

type User struct {
	ID int
	Username string
	Password string
	Email string
}


func (user *User)AddUser() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("AddUser预编译异常")
		return err
	}
	_, err = inStmt.Exec("admin", "123456", "123456@163.com")
	if err != nil {
		fmt.Println("AddUser执行异常")
		return err
	}
	return nil
}

func (user *User)AddUser2() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	_, err := utils.Db.Exec(sqlStr,"admin2", "123456", "123456@163.com")
	if err != nil {
		fmt.Println("AddUser2 执行异常")
		return err
	}
	return nil
}

func (user *User) GetUserById(userId int) (*User, error) {
	sqlStr := "select id, username, password, email from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, userId)
	var username string
	var password string
	var email string
	err := row.Scan(&username, &password, &email)

	if err != nil {
		return nil, err
	}
	u := &User {
		ID: userId,
		Username: username,
		Password: password,
		Email: email,
	}
	return u,nil
}

func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select id, username, password, email from users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*User

	for rows.Next() {
		var userId int
		var username string
		var password string
		var email string
		err := rows.Scan(&userId, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User {
			ID: userId,
			Username: username,
			Password: password,
			Email: email,
		}
		users = append(users, u)
	}

	return users,nil
}