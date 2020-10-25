package identity

import (
	"autodesk-go/utils"

	"github.com/labstack/echo/v4"
)

func SignupIdentity(c echo.Context) error {
	request := &UserRegisterRequest{}
	err := utils.ParseAndValidate(c, request)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	createErr, newIdentity := CreateNewIdentity(request)
	if createErr != nil {
		return utils.ErrorResponse(c, createErr)
	}

	return utils.SuccessResponse(c, newIdentity)
}

func GetUsername(c echo.Context) error {
	request := &UsernameFindRequest{}
	err := utils.ParseAndValidate(c, request)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	findErr, identity := FindByUsername(request.Username)
	if findErr != nil {
		return utils.ErrorResponse(c, findErr)
	}

	return utils.SuccessResponse(c, identity)
}

func VerifyIdentity(c echo.Context) error {
	request := &IdentityVerifyRequest{}
	err := utils.ParseAndValidate(c, request)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	verifyErr, identity := VerifyIdentityCredentials(request)
	if verifyErr != nil {
		return utils.ErrorResponse(c, verifyErr)
	}

	return utils.SuccessResponse(c, identity)
}

func RegisterRoute(e *echo.Echo) {
	e.POST("/identity/signup", SignupIdentity)
	e.POST("/identity/username/find", GetUsername)
	e.POST("/identity/signin", VerifyIdentity)
}
