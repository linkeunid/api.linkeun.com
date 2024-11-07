package models

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255)" validate:"required,min=2,max=255"`
	Username  string    `json:"username" gorm:"type:varchar(39);unique" validate:"required,min=2,max=39"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique" validate:"required,email"`
	Password  string    `json:"-" gorm:"type:varchar(255)" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()"`
}

type CreateUserRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=255"`
	Username        string `json:"username" validate:"required,min=2,max=39"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,eqfield=Password"`
}

type UpdateUserRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=255"`
	Username        string `json:"username" validate:"required,min=2,max=39"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignInResponse struct {
	User        *User  `json:"user"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
