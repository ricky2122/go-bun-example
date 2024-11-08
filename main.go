package main

import (
	"context"
	"errors"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

type UserID int

type User struct {
	ID       UserID
	Name     string
	Password string
	Email    string
	BirthDay string
}

func main() {
	conf := DBConfig{
		Host:     "localhost",
		Port:     "15432",
		DBName:   "bun_example",
		User:     "root",
		Password: "password",
	}

	db := NewDB(conf)
	ctx := context.TODO()

	// get users
	users, err := GetUsers(ctx, db)
	if err != nil {
		fmt.Printf("Failed getting users: %v", err)
		return
	}
	fmt.Printf("users: %v", users)

	// get user by id
	userID := 1
	user, err := GetUserByID(ctx, db, int64(userID))
	if err != nil {
		fmt.Printf("Failed getting user(id: %d): %v", userID, err)
		return
	}
	fmt.Printf("user(id:%d): %v", userID, *user)

	// create user
	newUser := User{
		Name:     "user04",
		Password: "example04",
		Email:    "example04@example.com",
		BirthDay: "2003-01-01",
	}
	createdUser, err := CreateUser(ctx, db, newUser)
	if err != nil {
		var appErr *CustomError
		if !errors.As(err, &appErr) {
			fmt.Printf("Failed creating user: %v", err)
			return
		}

		switch appErr.ErrCode {
		case DuplicateKeyErr:
			fmt.Printf("Failed creating user: %v", appErr)
			return
		}

	}
	fmt.Printf("created User: %v", createdUser)

	// delete user
	deleteUserID := 4
	if err := DeleteUser(ctx, db, UserID(deleteUserID)); err != nil {
		fmt.Printf("Failed delete user: %v", err)
	}
}
