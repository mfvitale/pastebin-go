package model

type Account int

const (
	Normal Account = iota
	Pro
)

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
