package app

import (
	// "github.com/nu7hatch/gouuid"
    "time"
    "golang.org/x/crypto/bcrypt"
)

/*
func generateId() string {
	id, _ := uuid.NewV4()
	return id.String()
}
*/

type User struct {
	Id        int    `json:"_" gorm:"primary_key"`
	Guid      string `json:"id" sql:"index"`
	Email     string `json:"email" sql:"index"`
	Username  string `json:"username"`
	Password  string `json:"_"`
    CreatedAt time.Time
}

func NewUser(email string, username string) *User {
    return &User{Guid: generateId(), Email: email, Username: username}
}

func (user *User) getToken() string {
    return CreateSignedValue(SECRET, AUTH_NAME, user.Guid, nil)
}

func (user *User) setPassword(plaintext string) error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashed)
    return nil
}

func (user *User) validate(plaintext string) bool {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plaintext)) == nil
}
