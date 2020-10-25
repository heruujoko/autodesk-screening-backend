package identity

import (
	"github.com/labstack/echo/v4"
	"autodesk-go/utils"
)

func SignupIdentity(c echo.Context) error {
	request := &UserRegisterRequest{}
	err := utils.ParseAndValidate(c, request)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}
	
	return utils.SuccessResponse(c, "SIGNUP");
}

func RegisterRoute(e *echo.Echo) {
	e.POST("/identity/signup", SignupIdentity)
}