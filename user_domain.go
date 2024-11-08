package main

import (
	"fmt"
	"time"
)

const BirthDayLayout = "2006-01-02"

type (
	UserID   int
	BirthDay time.Time
)

func (b BirthDay) String() string {
	return time.Time(b).Format(BirthDayLayout)
}

type User struct {
	id       UserID
	name     string
	password string
	email    string
	birthDay BirthDay
}

func NewUser(name, password, email, birthDay string) (*User, error) {
	parseBirthDay, err := time.Parse(BirthDayLayout, birthDay)
	if err != nil {
		return nil, err
	}
	return &User{
		name:     name,
		password: password,
		email:    email,
		birthDay: BirthDay(parseBirthDay),
	}, nil
}

func (u User) String() string {
	return fmt.Sprintf(
		"{id:%d name:%s password:%s email:%s birthDay:%s}",
		u.id, u.name, u.password, u.email, u.birthDay.String(),
	)
}

func (u *User) GetID() UserID {
	return u.id
}

func (u *User) SetID(id int) {
	u.id = UserID(id)
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetBirthDay() BirthDay {
	return u.birthDay
}

func (u *User) SetBirthDay(birthDay string) error {
	parseBirthDay, err := time.Parse(BirthDayLayout, birthDay)
	if err != nil {
		return err
	}

	u.birthDay = BirthDay(parseBirthDay)

	return nil
}
