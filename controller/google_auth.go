package controller

import (
	"fmt"
	"net/http"
	"user/config"
	"user/model/request"
	"user/model/response"
	"user/utils"

	error_handling "user/error"

	"golang.org/x/oauth2"
)

var Oauth2Config oauth2.Config

func (c *UserController) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	authURL := "https://accounts.google.com/o/oauth2/v2/auth?scope=https://www.googleapis.com/auth/userinfo.profile&access_type=offline&include_granted_scopes=true&response_type=code&state=state_parameter_passthrough_value&redirect_uri=" + config.ConfigVal.GooglAuth.RedirectURL + "&client_id=" + config.ConfigVal.GooglAuth.ClientID + ""
	googleAuthURL := response.GoogleAuthURL{AuthURL: authURL}
	utils.SuccessMessageResponse(w, 200, googleAuthURL)
}

func (c *UserController) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	googleAccessTokenRequest := request.GoogleAccessTokenRequest{
		Code:         code,
		ClientID:     config.ConfigVal.GooglAuth.ClientID,
		CLientSecret: config.ConfigVal.GooglAuth.ClientSecret,
		RedirectURI:  config.ConfigVal.GooglAuth.RedirectURL,
		GrantType:    "authorization_code",
	}
	var googleAccessTokenResponse response.GoogleAccessTokenResponse
	err := utils.ExternalURLCall("POST", "https://oauth2.googleapis.com/token", googleAccessTokenRequest, googleAccessTokenResponse)
	if err != nil {
		fmt.Println("sdf")
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	var googleUserInfo response.GoogleUserInfo
	err = utils.ExternalURLCall("GET", "https://www.googleapis.com/oauth2/v3/userinfo?access_token="+googleAccessTokenResponse.AccessToken, nil, googleUserInfo)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		error_handling.ErrorMessageResponse(w, error_handling.OTPGenerateError)
		return
	}
	hashedOTP, err := utils.Bcrypt(otp)
	if err != nil {
		error_handling.ErrorMessageResponse(w, error_handling.BcryptError)
		return
	}
	signup := request.Signup{
		FirstName: googleUserInfo.GivenName,
		LastName:  googleUserInfo.FamilyName,
		Email:     googleUserInfo.Email,
		EventType: "google",
	}
	err = c.repo.StoreOTP(signup, hashedOTP, "signup")
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	subject := "OTP for login/signup:"
	err = utils.SendOTPEmail(signup.Email, otp, subject)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
}
