package model

import "github.com/mfvitale/pastebin-go/internal/client/dto"

type Account int

const (
	Normal Account = iota
	Pro
)

func NewAccount(account int) Account {

	switch account {
	case 0:
		return Normal
	case 1:
		return Pro
	default:
		return -1
	}
}

type User struct {
	Name       string
	Format     string
	Expiration string
	Avatar     string
	Visibility Visibility
	Website    string
	Email      string
	Location   string
	Type       Account
}

func NewUser(user dto.User) *User {
	return &User{user.Name,
		user.Format,
		user.Expiration,
		user.Avatar,
		NewVisibility(user.Visibility),
		user.Website,
		user.Email,
		user.Location,
		NewAccount(user.Type),
	}
}
