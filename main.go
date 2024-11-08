package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
)

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
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

	// create article table
	if err := CreateArticleTable(ctx, db); err != nil {
		fmt.Printf("Failed creating article table: %v", err)
		return
	}

	// truncate article table
	if err := TruncateArticleTable(ctx, db); err != nil {
		fmt.Printf("Failed truncating article table: %v", err)
		return
	}

	// drop article table
	if err := DropArticleTable(ctx, db); err != nil {
		fmt.Printf("Failed deleting article table: %v", err)
		return
	}

	// userDBExecution(ctx, db)
}

func userDBExecution(ctx context.Context, db *bun.DB) {
	// get users
	users, err := GetUsers(ctx, db)
	if err != nil {
		fmt.Printf("Failed getting users: %v\n", err)
		return
	}
	fmt.Print("Get users:")
	for _, user := range users {
		fmt.Printf(" %s", user)
	}
	fmt.Println()

	// get user by id
	userID := 1
	user, err := GetUserByID(ctx, db, int64(userID))
	if err != nil {
		fmt.Printf("Failed getting user(id: %d): %v\n", userID, err)
		return
	}
	fmt.Printf("Get user: %s\n", user)

	// create user
	createUser, _ := NewUser("user04", "example04", "example04@example.com", "2004-01-01")
	createdUser, err := CreateUser(ctx, db, *createUser)
	if err != nil {
		checkCustomError(err)
	}
	fmt.Printf("Created User: %s\n", createdUser)

	// update user
	updateUser := createdUser
	updateUser.SetName("update_user04")
	updateUser.SetPassword("update_example04")
	updateUser.SetEmail("update_example04@example.com")
	_ = updateUser.SetBirthDay("2004-02-02")
	updatedUser, err := UpdateUser(ctx, db, *updateUser)
	if err != nil {
		checkCustomError(err)
	}
	fmt.Printf("Updated user: %s\n", updatedUser)

	// delete user
	deleteUserID := createdUser.GetID()
	if err := DeleteUser(ctx, db, UserID(deleteUserID)); err != nil {
		fmt.Printf("Failed deleting user: %v\n", err)
		return
	}
	fmt.Printf("Delete userID: %d\n", deleteUserID)

	// bulk insert users
	createUsers := make([]User, 0, 2)
	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("user0%d", i+5)
		password := fmt.Sprintf("example0%d", i+5)
		email := fmt.Sprintf("example0%d@example.com", i+5)
		birthDay := fmt.Sprintf("200%d-01-01", i+5)
		createUser, _ = NewUser(name, password, email, birthDay)
		createUsers = append(createUsers, *createUser)
	}
	createdUsers, err := BulkInsertUsers(ctx, db, createUsers)
	if err != nil {
		checkCustomError(err)
	}
	fmt.Printf("created users: %+v\n", createdUsers)

	// bulk update users
	updateUsers := make([]User, 0, len(createdUsers))
	for i, createdUser := range createdUsers {
		updateUser := createdUser
		name := fmt.Sprintf("update_user0%d", i+5)
		password := fmt.Sprintf("update_example0%d", i+5)
		email := fmt.Sprintf("update_example0%d@example.com", i+5)
		birthDay := fmt.Sprintf("200%d-02-02", i+5)

		updateUser.SetName(name)
		updateUser.SetPassword(password)
		updateUser.SetEmail(email)
		_ = updateUser.SetBirthDay(birthDay)

		updateUsers = append(updateUsers, updateUser)
	}
	updatedUsers, err := BulkUpdateUsers(ctx, db, updateUsers)
	if err != nil {
		checkCustomError(err)
	}
	fmt.Printf("Update users: %+v", updatedUsers)

	// bulk delete users
	deletedUserIDs := make([]UserID, 0, len(createdUsers))
	for _, createdUser := range createdUsers {
		deletedUserIDs = append(deletedUserIDs, createdUser.GetID())
	}
	if err := BulkDeleteUsers(ctx, db, deletedUserIDs); err != nil {
		fmt.Printf("Failed deleting users: %v", err)
		return
	}
	fmt.Printf("Delete userIDs: %+v", deletedUserIDs)
}

func checkCustomError(err error) {
	var customErr *CustomError
	if !errors.As(err, &customErr) {
		fmt.Printf("Failed updating users: %v\n", err)
		return
	}

	switch customErr.ErrCode {
	case ErrDuplicateKey:
		fmt.Printf("Failed updating users: %v\n", customErr)
		return
	}
}
