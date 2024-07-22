package factory

import (
	"example/template/internal/domain/entity"
)

// User entity might be more complicated, and consist
// nested object, so factory can help create new object
type User struct{}

func (u *User) Generate(id string, name string, email string) entity.User {
	return entity.User{
		Id:    id,
		Name:  name,
		Email: email,
	}
}
