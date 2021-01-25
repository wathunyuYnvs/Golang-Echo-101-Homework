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

// @Summary Get profile
// @Description Get user information
// @Tags Private
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer sfds1f32ds....fsd"
// @Accept json
// @Produce json
// @Success 200 {object} user.ProfileResponseBody
// @Failure 417 {string} string
// @Failure 400 {string} string
// @Router /private/users/profile [Get]
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

// @Summary Edit Profile
// @Description Edit user information
// @Tags Private
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer sfds1f32ds....fsd"
// @Accept json
// @Produce json
// @Param edit user body user.UserRequestBody true "Edit user"
// @Success 200 {object} user.EditProfileResponseBody
// @Failure 417 {string} string
// @Failure 400 {string} string
// @Router /private/users/profile [Put]
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

// @Summary Change password
// @Description Edit password
// @Tags Private
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer sfds1f32ds....fsd"
// @Accept json
// @Produce json
// @Param change password body user.ChangePaswordRequestBody true "Edit password"
// @Success 200 {object} user.EditProfileResponseBody
// @Failure 417 {string} string
// @Failure 400 {string} string
// @Router /private/users/changePassword [Put]
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
