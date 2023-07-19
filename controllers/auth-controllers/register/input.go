package register

import "mime/multipart"

type RegisterInput struct {
	Username string                `form:"username" validate:"required,lowercase"`
	Email    string                `form:"email" validate:"required,email"`
	Password string                `form:"password" validate:"required,gte=8"`
	Image    *multipart.FileHeader `form:"image"`
}
