package identity

import (
	"autodesk-go/utils"
)

type Identity struct {
	utils.ModelBase
	Email        string `json:"email" gorm:"unique;not null";sql:"index"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username" gorm:"unique;not null";sql:"index"`
	PasswordHash string `json:"-"`
}

type UserRegisterRequest struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Username  string `json:"username" validate:"required"`
}

type UsernameFindRequest struct {
	Username string `json:"username" validate:"required"`
}

type IdentityVerifyRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
