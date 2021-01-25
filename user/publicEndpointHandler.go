package user

import (
	"myecho/db"
	"myecho/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PublicEndpointHadlerInit(r *echo.Group, resource *db.Resource) {

	NewUserService(resource)

	r.POST("/users/login", login())
	r.POST("/users/register", register())
}

// @Summary Login
// @Description Login
// @Tags Public
// @Accept json
// @Produce json
// @Param login body user.UserLoginRequestBody true "Logged in"
// @Success 200 {string} string
// @Failure 417 {string} string
// @Router /public/users/login [post]
func login() func(c echo.Context) error {
	return func(c echo.Context) error {
		newUser := UserLoginRequestBody{}
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		userDetail, err := LoginService(newUser)
		if err != nil {
			return c.String(http.StatusExpectationFailed, err.Error())
		}
		t, _ := middlewares.Sign(middlewares.Payload{ID: userDetail.ID, Name: userDetail.Name})
		return c.String(http.StatusOK, t)
	}
}

// @Summary Register
// @Description Register
// @Tags Public
// @Accept json
// @Produce json
// @Param register body user.UserRegisterRequestBody true "Register"
// @Success 201 {object} user.UserLoginResponseBody
// @Failure 417 {string} string
// @Router /public/users/register [post]
func register() func(c echo.Context) error {
	return func(c echo.Context) error {
		newUser := UserRegisterRequestBody{}
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		userDetail, err := RegisterService(newUser)
		if err != nil {
			return c.String(http.StatusExpectationFailed, err.Error())
		}
		t, _ := middlewares.Sign(middlewares.Payload{ID: userDetail.ID, Name: userDetail.Name})
		return c.JSON(http.StatusCreated, t)
	}
}
