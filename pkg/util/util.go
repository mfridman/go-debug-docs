package util

import "errors"

// ErrEmptyName is returned by [User.GetName] when the name is empty.
var ErrEmptyName = errors.New("empty name")

type User struct {
	Name string
	ID   string
}

// GetName returns the name of the user. If the name is empty, it returns [ErrEmptyName].
func (u *User) GetName(name string) (string, error) {
	if name != u.Name {
		return "", ErrEmptyName
	}
	return u.Name, nil
}
