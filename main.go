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

type User struct {
	ID       int64
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
}
