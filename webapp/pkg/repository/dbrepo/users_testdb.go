package dbrepo

import (
	"database/sql"
	"errors"
	"time"
	"webapp/pkg/data"
)

type TestDBRepo struct{}

func (t *TestDBRepo) Connection() *sql.DB {
	return nil
}

func (t *TestDBRepo) AllUsers() ([]*data.User, error) {
	var users []*data.User
	return users, nil
}
func (t *TestDBRepo) GetUser(id int) (*data.User, error) {
	var user = data.User{}
	if id == 1 {
		user = data.User{
			ID:        1,
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@example.com",
			Password:  "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK",
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
func (t *TestDBRepo) GetUserByEmail(email string) (*data.User, error) {
	if email == "admin@example.com" {
		user := data.User{
			ID:        1,
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@example.com",
			Password:  "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK",
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		return &user, nil
	}

	return nil, errors.New("user not found")
}
func (t *TestDBRepo) UpdateUser(u data.User) error {
	if u.ID == 1 {
		return nil
	}
	return errors.New("update failed - no user found")
}
func (t *TestDBRepo) DeleteUser(id int) error {
	return nil
}
func (t *TestDBRepo) InsertUser(user data.User) (int, error) {
	return 2, nil
}
func (t *TestDBRepo) ResetPassword(id int, password string) error {
	return nil
}
func (t *TestDBRepo) InsertUserImage(i data.UserImage) (int, error) {
	return 1, nil
}
