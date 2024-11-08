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
		fmt.Printf("Failed getting users: %v\n", err)
		return
	}
	fmt.Printf("users: %v\n", users)

	// get user by id
	userID := 1
	user, err := GetUserByID(ctx, db, int64(userID))
	if err != nil {
		fmt.Printf("Failed getting user(id: %d): %v\n", userID, err)
		return
	}
	fmt.Printf("user(id:%d): %v\n", userID, *user)

	// create user
	newUser := User{
		Name:     "user04",
		Password: "example04",
		Email:    "example04@example.com",
		BirthDay: "2003-01-01",
	}
	createdUser, err := CreateUser(ctx, db, newUser)
	if err != nil {
		var customErr *CustomError
		if !errors.As(err, &customErr) {
			fmt.Printf("Failed creating user: %v\n", err)
			return
		}

		switch customErr.ErrCode {
		case DuplicateKeyErr:
			fmt.Printf("Failed creating user: %v\n", customErr)
			return
		}

	}
	fmt.Printf("created User: %v\n", createdUser)

	// update user
	updateUser := User{
		ID:       createdUser.ID,
		Name:     "user05",
		Password: "example05",
		Email:    "example05@example.com",
		BirthDay: "2004-01-01",
	}
	updatedUser, err := UpdateUser(ctx, db, updateUser)
	if err != nil {
		var customErr *CustomError
		if !errors.As(err, &customErr) {
			fmt.Printf("Failed updating user: %v\n", err)
			return
		}

		switch customErr.ErrCode {
		case DuplicateKeyErr:
			fmt.Printf("Failed updating user: %v\n", customErr)
			return
		}
	}
	fmt.Printf("updated user: %v\n", updatedUser)

	// delete user
	deleteUserID := createdUser.ID
	if err := DeleteUser(ctx, db, UserID(deleteUserID)); err != nil {
		fmt.Printf("Failed deleting user: %v\n", err)
		return
	}
}
