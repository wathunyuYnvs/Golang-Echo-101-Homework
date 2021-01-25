package user

type UserLoginResponseBody struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type ProfileResponseBody struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
type EditProfileResponseBody struct {
	result bool `json:"result"`
}
