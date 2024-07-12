package jsonformat

import "github.com/loyalsfc/investrite/models"

type SigninStruct struct {
	AccessToken string      `json:"access_token"`
	UserData    models.User `json:"user_data"`
}

func SignInToSignIn(user models.User, token string) SigninStruct {
	return SigninStruct{
		AccessToken: token,
		UserData:    user,
	}
}
