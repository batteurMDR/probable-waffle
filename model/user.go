package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// User represents an administrator of the service.
type User struct {
	// ID is the unique identifier of the user in the database.
	ID string `json:"id"`
	// FirstName stands for the firstname of the user.
	FirstName string `json:"first_name"`
	// LastName stands for the lastname of the user.
	LastName string `json:"last_name"`
	// Email is the email used for login and communicating with the user.
	Email string `json:"email" gorm:"index:,unique"`
	// Password is the password hashed in sha256 of the user.
	Password Password `json:"password"`
}

// UnmarshalJSON implement the Umarshaler contract interface on the user.
func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}

	return json.Marshal(aux)
}

func hashPassword(value string) (res Password) {
	h := sha256.New()
	h.Write([]byte(value))
	res = Password(fmt.Sprintf("%x", h.Sum(nil)))
	return res
}

// Password is a type for containing password.
type Password string

// MarshalJSON implement the Marshaler contract interface on the Password.
func (p *Password) UnmarshalJSON(data []byte) error {

	var s string = ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*p = hashPassword(s)

	return nil
}

type LoginPayload struct {
	Email string `json:"email"`
	// Password is the password hashed in sha256 of the user.
	Password Password `json:"password"`
}
