package dto

type CreateUser struct {
	Name  string
	Email string
}

type UserFilter struct {
	ID    string
	Name  string
	Email string
}
