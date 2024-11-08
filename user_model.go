package main

import (
	"context"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int    `bun:"id,pk,autoincrement"`
	Name     string `bun:"name,notnull,unique"`
	Password string `bun:"password,notnull"`
	Email    string `bun:"email,notnull,unique"`
	BirthDay string `bun:"birth_day,notnull"`
}

func GetUserByID(ctx context.Context, db *bun.DB, userID int64) (*User, error) {
	var userModel UserModel
	if err := db.NewSelect().Model(&userModel).Where("id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}

	user := convertToUser(userModel)

	return &user, nil
}

func GetUsers(ctx context.Context, db *bun.DB) ([]User, error) {
	var userModels []UserModel
	if err := db.NewSelect().Model(&userModels).Scan(ctx); err != nil {
		return nil, err
	}

	users := convertToUsers(userModels)

	return users, nil
}

func convertToUser(userModel UserModel) User {
	return User{
		ID:       UserID(userModel.ID),
		Name:     userModel.Name,
		Password: userModel.Password,
		Email:    userModel.Email,
		BirthDay: userModel.BirthDay,
	}
}

func convertToUsers(userModels []UserModel) []User {
	users := make([]User, 0, len(userModels))

	for _, userModel := range userModels {
		users = append(users, convertToUser(userModel))
	}

	return users
}
