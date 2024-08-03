package product

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Entity struct {
	ID          string  `json:"id,omitempty" db:"id"`
	Title       string  `json:"title,omitempty" db:"title"`
	Description string  `json:"description,omitempty" db:"description"`
	Price       float64 `json:"price,omitempty" db:"price"`
	Category    string  `json:"category,omitempty" db:"category"`
	Quantity    int64   `json:"quantity,omitempty" db:"quantity"`
}

func (req *Entity) Validate() error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Title, validation.Required, validation.Length(5, 50)),
		validation.Field(&req.Description, validation.Required, validation.Length(15, 1500)),
		validation.Field(&req.Price, validation.Required, validation.Min(0.1)),
		validation.Field(&req.Category, validation.Required, validation.Length(5, 50)),
		validation.Field(&req.Quantity, validation.Required, validation.Min(1)),
	)
}
