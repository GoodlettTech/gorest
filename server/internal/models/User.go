package Models

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=4,max=16"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type Credentials struct {
	Username string `json:"username" validate:"required,min=4,max=16"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
