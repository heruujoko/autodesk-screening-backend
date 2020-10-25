package identity

import (
	"autodesk-go/utils"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

var once sync.Once
var DB *gorm.DB

func init() {
	once.Do(func() {
		DB = utils.GetInstanceDB()
		fmt.Println("init identity service")
		DB.AutoMigrate(&Identity{})
	})
}

func CreateNewIdentity(data *UserRegisterRequest) (*utils.ErrorBag, *Identity) {
	hashedPassword, errHashing := utils.HashString(data.Password)
	if errHashing != nil {
		errors := utils.ErrorBag{
			Code:    "UNABLE_TO_REGISTER",
			Message: errHashing.Error(),
			Data:    errHashing,
		}
		return &errors, nil
	}

	identity := Identity{
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		PasswordHash: hashedPassword,
		Username:     data.Username,
	}
	result := DB.Save(&identity)
	if result.Error != nil {
		errors := utils.ErrorBag{
			Code:    "UNABLE_TO_REGISTER",
			Message: result.Error.Error(),
			Data:    result.Error.Error(),
		}
		return &errors, nil
	} else {
		return nil, nil
	}
}

func FindByUsername(username string) (*utils.ErrorBag, *Identity) {
	user := &Identity{}
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		respObject := utils.FieldError{
			FieldName: "username",
			Status:    "not found",
			Type:      "string",
			Message:   "The username is not recognized",
		}
		errors := utils.ErrorBag{
			Code:    "NOT_FOUND",
			Message: "The username is not recognized",
			Data:    []utils.FieldError{respObject},
		}
		return &errors, nil
	} else {
		if len(user.FirstName) == 0 {
			return nil, nil
		}
		return nil, user
	}
}

func VerifyIdentityCredentials(data *IdentityVerifyRequest) (*utils.ErrorBag, *Identity) {
	errFind, currentIdentity := FindByUsername(data.Username)
	if errFind != nil {
		return errFind, nil
	}

	passMatch := utils.CheckPasswordHash(data.Password, currentIdentity.PasswordHash)
	if passMatch {
		return nil, currentIdentity
	} else {
		errors := utils.ErrorBag{
			Code:    "AUTENTICTION_FAILED",
			Message: "password not match",
			Data:    "password not match",
		}
		return &errors, nil
	}
}
