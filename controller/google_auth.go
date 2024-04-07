package controller

import (
	"fmt"
	"net/http"
	"user/config"
	"user/constant"
	"user/model/request"
	"user/model/response"
	"user/utils"

	error_handling "user/error"
)

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

	var bodyDataResponse map[string]interface{}

	bodyDataResponse, err := utils.ExternalURLCall("POST", "https://oauth2.googleapis.com/token", googleAccessTokenRequest, bodyDataResponse)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	accessToken := bodyDataResponse["access_token"]
	fmt.Println(accessToken)
	bodyDataResponse, err = utils.ExternalURLCall("GET", "https://www.googleapis.com/oauth2/v3/userinfo?access_token="+accessToken.(string), nil, bodyDataResponse)
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
	storeOTP := request.StoreOTP{
		Email:     bodyDataResponse["email"].(string),
		EventType: "google_login",
		LoginType: "google_login",
		HashedOTP: hashedOTP,
	}
	err = c.repo.StoreOTP(storeOTP)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	subject := "OTP for login/signup:"
	go utils.SendOTPEmail(bodyDataResponse["email"].(string), otp, subject)
	// if err != nil {
	// 	error_handling.ErrorMessageResponse(w, err)
	// 	return
	// }
	googleInfo := response.GoogleUserInfo{
		FirstName: bodyDataResponse["given_name"].(string),
		LastName:  bodyDataResponse["family_name"].(string),
		Email:     bodyDataResponse["email"].(string),
		Message:   constant.OTP_SENT,
	}
	utils.SuccessMessageResponse(w, 200, googleInfo)
}
