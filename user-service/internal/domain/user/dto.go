package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (req *Request) Validate() error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&req.Email, validation.Required, is.Email),
	)
}

type Response struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ParseFromEntity(user *Entity) *Response {
	res := &Response{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return res
}

func ParseFromEntities(users []*Entity) []*Response {
	res := make([]*Response, 0)

	for _, user := range users {
		res = append(res, ParseFromEntity(user))
	}

	return res
}
