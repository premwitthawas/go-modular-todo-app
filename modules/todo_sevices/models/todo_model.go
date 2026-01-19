package todoModels

import (
	"time"
)

type (
	TodoCreateReq struct {
		Title string `json:"title" form:"title" validate:"required"`
		Email string `json:"email" form:"email" validate:"required,email"`
	}
	TodoRes struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Email     string    `json:"email"`
		Done      bool      `json:"done"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
