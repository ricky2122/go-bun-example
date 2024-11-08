package main

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int    `bun:"id,pk,autoincrement"`
	Name     string `bun:"name,notnull,unique"`
	Password string `bun:"password,notnull"`
	Email    string `bun:"email,notnull,unique"`
	BirthDay string `bun:"birth_day,notnull"`
}

func CreateUser(ctx context.Context, db *bun.DB, newUser User) (*User, error) {
	userModel := convertToUserModel(newUser)
	_, err := db.NewInsert().
		Model(&userModel).
		Returning("id").
		Exec(ctx)
	if err != nil {
		// Check if the error is a pgdriver.Error
		var pgdErr pgdriver.Error
		if errors.As(err, &pgdErr) {
			// SQLState 23305 indicates a unique violation
			if pgdErr.Field('C') == "23505" {
				return nil, DuplicateKeyErr.Wrap(err, "duplicate key error")
			}
			return nil, pgdErr
		}
		return nil, err
	}

	createdUser := convertToUser(userModel)

	return &createdUser, nil
}

func BulkInsertUsers(ctx context.Context, db *bun.DB, createUsers []User) ([]User, error) {
	createUserModels := convertToUserModels(createUsers)

	if _, err := db.NewInsert().Model(&createUserModels).Exec(ctx); err != nil {
		return nil, err
	}

	createdUsers := convertToUsers(createUserModels)

	return createdUsers, nil
}

func DeleteUser(ctx context.Context, db *bun.DB, deleteUserID UserID) error {
	deleteUserModel := UserModel{ID: int(deleteUserID)}
	if _, err := db.NewDelete().Model(&deleteUserModel).WherePK().Exec(ctx); err != nil {
		return err
	}

	return nil
}

func BulkDeleteUsers(ctx context.Context, db *bun.DB, deleteUserIDs []UserID) error {
	deleteUserModels := make([]UserModel, 0, len(deleteUserIDs))

	for _, deleteUserID := range deleteUserIDs {
		deleteUserModels = append(deleteUserModels, UserModel{ID: int(deleteUserID)})
	}

	if _, err := db.NewDelete().Model(&deleteUserModels).WherePK().Exec(ctx); err != nil {
		return err
	}

	return nil
}

func UpdateUser(ctx context.Context, db *bun.DB, updateUser User) (*User, error) {
	updateUserModel := convertToUserModel(updateUser)
	if _, err := db.NewUpdate().Model(&updateUserModel).WherePK().Exec(ctx); err != nil {
		return nil, err
	}

	return &updateUser, nil
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

func convertToUserModel(user User) UserModel {
	return UserModel{
		ID:       int(user.ID),
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		BirthDay: user.BirthDay,
	}
}

func convertToUserModels(users []User) []UserModel {
	userModels := make([]UserModel, 0, len(users))

	for _, user := range users {
		userModels = append(userModels, convertToUserModel(user))
	}

	return userModels
}
