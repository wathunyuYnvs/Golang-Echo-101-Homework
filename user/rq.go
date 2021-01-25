package user

type UserRequestBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserLoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterRequestBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Name            string `json:"name"`
	Age             int    `json:"age"`
}

type ChangePaswordRequestBody struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
