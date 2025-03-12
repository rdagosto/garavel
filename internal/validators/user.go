package validators

type UserStore struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdate struct {
	Password string `json:"password" validate:"required,min=6"`
}
