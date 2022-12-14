package local

import "time"

type User struct {
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type PublicUser struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type SimpleUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

var users = "users"

func (m *Manager) CreateUser(u SimpleUser) (PublicUser, error) {
	user := User{
		Username: u.Username,
		Password: u.Password,
		FullName: u.FullName,
		Email:    u.Email,
	}
	err := m.handler.Table(users).Create(&user).Error
	return PublicUser{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}, err
}

func (m *Manager) GetUser(username string) (User, error) {
	var u User
	err := m.handler.Table(users).Where("username = ?", username).Take(&u).Error
	return u, err
}
