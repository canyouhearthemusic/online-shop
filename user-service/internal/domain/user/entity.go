package user

type Entity struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserError struct {
	message string
}

func (e *UserError) Error() string {
	return e.message
}

var (
	ErrExists     = &UserError{"User already exists"}
	ErrNotFound   = &UserError{"User not found"}
	ErrSearch     = &UserError{"User search error"}
	ErrBadRequest = &UserError{"User bad request"}
)
