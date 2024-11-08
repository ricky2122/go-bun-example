package main

import (
	"context"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name,notnull,unique"`
	Password string `bun:"password,notnull"`
	Email    string `bun:"email,notnull,unique"`
	BirthDay string `bun:"birth_day,notnull"`
}

func GetUsers(ctx context.Context, db *bun.DB) ([]User, error) {
	var userModels []UserModel
	if err := db.NewSelect().Model(&userModels).Scan(ctx); err != nil {
		return nil, err
	}

	users := convertToUsers(userModels)

	return users, nil
}

func convertToUsers(userModels []UserModel) []User {
	users := make([]User, len(userModels))

	for _, userModel := range userModels {
		user := User{
			ID:       userModel.ID,
			Name:     userModel.Name,
			Password: userModel.Password,
			Email:    userModel.Email,
			BirthDay: userModel.BirthDay,
		}
		users = append(users, user)
	}

	return users
}
