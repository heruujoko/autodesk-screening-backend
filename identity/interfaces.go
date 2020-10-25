package identity

import (
	"autodesk-go/utils"
)

type Identity struct {
	utils.ModelBase
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Username     string `json:"username" gorm:"unique;not null";sql:"index"`
	PasswordHash string `json:"-"`
}

type UserRegisterRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
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
