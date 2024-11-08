package main

import (
	"context"
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
	}
	fmt.Printf("user(id:%d): %v", userID, *user)
}
