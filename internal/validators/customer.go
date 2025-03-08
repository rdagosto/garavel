package validators

type CustomerStore struct {
	Name  string `json:"name"  validate:"required|min=3"`
	Email string `json:"email" validate:"required,email"`
}

type CustomerUpdate struct {
	Name  string `json:"name"  validate:"min=3"`
	Email string `json:"email" validate:"email"`
}
