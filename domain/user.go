package domain

import "fmt"

type User struct {
	ID       string
	Username string
}

func (u *User) Validate() error {
	if u.ID == "" {
		return fmt.Errorf("ID can't be blank.")
	}

	if u.Username == "" {
		return fmt.Errorf("Username can't be blank.")
	}

	return nil
}
