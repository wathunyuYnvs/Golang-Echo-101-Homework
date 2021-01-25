package user

import (
	"myecho/db"
	"myecho/middlewares"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func PrivateEndpointHadlerInit(r *echo.Group, resource *db.Resource) {

	NewUserService(resource)

	r.GET("/users/profile", getProfile())
	r.PUT("/users/profile", editProfile())
	r.PUT("/users/changePassword", changePassword())
}

func getProfile() func(c echo.Context) error {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*middlewares.JwtCustomClaims)
		r, err := GetProfileSrvice(claims.ID)
		if err != nil {
			return c.String(http.StatusExpectationFailed, err.Error())
		}
		return c.JSON(http.StatusOK, r)
	}
}

func editProfile() func(c echo.Context) error {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*middlewares.JwtCustomClaims)

		body := UserRequestBody{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		r, err := EditProfileService(claims.ID, body)
		if err != nil {
			return c.String(http.StatusExpectationFailed, err.Error())
		}
		return c.JSON(http.StatusOK, r)
	}
}
func changePassword() func(c echo.Context) error {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*middlewares.JwtCustomClaims)

		body := ChangePaswordRequestBody{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		r, err := ChangePasswordService(claims.ID, body)
		if err != nil {
			return c.String(http.StatusExpectationFailed, err.Error())
		}
		return c.JSON(http.StatusOK, r)
	}
}
