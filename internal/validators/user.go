package validators

type UserStore struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required|min=6"`
}

type UserUpdate struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6"`
}
