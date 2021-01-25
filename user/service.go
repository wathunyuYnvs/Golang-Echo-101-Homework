package user

import (
	"errors"
	"fmt"
	"myecho/db"
	"myecho/helper"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userRepository Repository

// NewUserAPI to create the router of user
func NewUserService(resource *db.Resource) {
	// Create repository
	userRepository = NewUserRepository(resource)
}
func LoginService(body UserLoginRequestBody) (UserLoginResponseBody, error) {
	user, err := userRepository.FindLogin(body.Email)
	if err != nil {
		return UserLoginResponseBody{}, err
	}

	if user == nil {
		return UserLoginResponseBody{}, errors.New("User or Password not found!")
	}
	fmt.Println("Hi", helper.ComparePassword(user.Password, body.Password))
	if !helper.ComparePassword(user.Password, body.Password) {
		return UserLoginResponseBody{}, errors.New("User or Password not found!")
	}
	var r UserLoginResponseBody
	r.ID = user.Id.Hex()
	r.Name = user.Name
	return r, nil
}

func RegisterService(body UserRegisterRequestBody) (UserLoginResponseBody, error) {
	if body.Password != body.ConfirmPassword {
		return UserLoginResponseBody{}, errors.New("Invalid Password")
	}
	user, _ := userRepository.FindLogin(body.Email)
	if user != nil {
		return UserLoginResponseBody{}, errors.New("User exist!")
	}

	var newUser User
	newUser.Id = primitive.NewObjectID()
	newUser.Email = body.Email
	newUser.Name = body.Name
	newUser.Age = body.Age
	newUser.Password = helper.EncodePassword(body.Password)
	insertData, _ := userRepository.CreateOne(newUser)
	var r UserLoginResponseBody
	r.Name = insertData.Name
	r.ID = insertData.Id.Hex()
	return r, nil
}

func GetProfileSrvice(id string) (ProfileResponseBody, error) {
	user, err := userRepository.GetByID(id)
	if err != nil {
		return ProfileResponseBody{}, err
	}
	r := ProfileResponseBody{}
	r.Name = user.Name
	r.Email = user.Email
	r.Age = user.Age
	return r, nil
}

func EditProfileService(id string, body UserRequestBody) (EditProfileResponseBody, error) {
	user, err := userRepository.GetByID(id)
	if err != nil {
		return EditProfileResponseBody{}, errors.New("User not found")
	}

	user.Name = body.Name
	user.Age = body.Age

	if _, err := userRepository.UpdateOne(id, *user); err != nil {
		return EditProfileResponseBody{}, err
	}

	r := EditProfileResponseBody{}
	r.result = true
	return r, nil
}

func ChangePasswordService(id string, body ChangePaswordRequestBody) (EditProfileResponseBody, error) {
	if body.NewPassword != body.ConfirmPassword {
		return EditProfileResponseBody{}, errors.New("Confirm Password not match")
	}
	user, err := userRepository.GetByID(id)
	if err != nil {
		return EditProfileResponseBody{}, errors.New("User not found")
	}
	if !helper.ComparePassword(user.Password, body.OldPassword) {
		return EditProfileResponseBody{}, errors.New("Old password not match!")
	}
	user.Password = helper.EncodePassword(body.NewPassword)

	_, e := userRepository.UpdateOne(id, *user)
	if e != nil {
		return EditProfileResponseBody{}, err
	}
	return EditProfileResponseBody{result: true}, nil
}
