package error

import (
	"encoding/json"
	"net/http"

	"github.com/gookit/validate"
	_ "github.com/lib/pq"
)

func init() {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})
	// validate.AddValidator("emailOrPhoneNumber", func(val any) bool {
	// 	// do validate val ...

	// 	return true
	// })
	validate.AddGlobalMessages(map[string]string{
		// "minLength": "OO! {field} min length is %d",
		// "required": "oh! the {field} is required",
		// "email": "email is invalid",
		"emailOrPhoneNumber": "phoneNumber and countryCode must be null with the email",
	  })
}

// func (f UserForm) Messages() map[string]string {
// 	return validate.MS{
// 		"required": "oh! the {field} is required",
// 		"email": "email is invalid",
// 		"Name.required": "message for special field",
// 		"Age.int": "age must int",
// 		"Age.min": "age min value is 1",
// 	}
// }

type CustomError struct {
	StatusCode   int           `json:"statusCode" validate:"required" `
	ErrorMessage string        `json:"errorMessage" validate:"required" `
	InvalidData  []InvalidData `json:"invalidData" validate:"omitempty" `
}

type InvalidData struct {
	Field string
	Error map[string]string
}

func (c CustomError) Error() string {
	return c.ErrorMessage
}

func ErrorMessageResponse(w http.ResponseWriter, err error) {
	if error, ok := err.(CustomError); ok {
		response, _ := json.MarshalIndent(error, "", "  ")
		w.WriteHeader(error.StatusCode)
		w.Write(response)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	response, _ := json.MarshalIndent(InternalServerError, "", "  ")
	w.Write(response)
}

func CreateCustomError(errorMessage string, statusCode int, invalidData ...InvalidData) error {
	return CustomError{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
		InvalidData:  invalidData,
	}
}

var (
	UnmarshalError         = CreateCustomError("Error while unmarshling data.", http.StatusUnauthorized)
	InternalServerError    = CreateCustomError("Internal Server Error.", http.StatusInternalServerError)
	OTPGenerateError       = CreateCustomError("Error at generating OTP.", http.StatusInternalServerError)
	BcryptError            = CreateCustomError("Error at bcypting.", http.StatusInternalServerError)
	SendEmailError         = CreateCustomError("Error at sending email.", http.StatusInternalServerError)
	SendMessageError       = CreateCustomError("Error at sending message.", http.StatusInternalServerError)
	ExpiredOTP             = CreateCustomError("OTP expired.", http.StatusGone)
	InvalidOTP             = CreateCustomError("OTP is invalid.", http.StatusBadRequest)
	UserAlreadyExist       = CreateCustomError("User already exist.", http.StatusBadRequest)
	UserDoesNotExist       = CreateCustomError("User does not exist.", http.StatusNotFound)
	HeaderDataMissing      = CreateCustomError("Required header not found.", http.StatusBadRequest)
	UsernameAlreadyTaken   = CreateCustomError("Username already taken.", http.StatusBadRequest)
	InvalidDetails         = CreateCustomError("Invalid details provided.", http.StatusBadRequest)
	AlreadyBlocked         = CreateCustomError("Already blocked.", http.StatusBadRequest)
	JWTErrSignatureInvalid = CreateCustomError("Invalid signature on jwt token.", http.StatusUnauthorized)
	JWTTokenInvalid        = CreateCustomError("Invalid jwt token.", http.StatusBadRequest)
	JWTTokenInvalidDetails = CreateCustomError("Invalid jwt token details.", http.StatusBadRequest)
)

// func DatabaseError(err error) error {
// 	if dbErr, ok := err.(*pq.Error); ok {
// 		errCode := dbErr.Code
// 		switch errCode {
// 		case "23502":
// 			// not-null constraint violation
// 			return CreateCustomError("Some required data was left out", 400)

// 		case "23503":
// 			// foreign key violation
// 			return CreateCustomError("This record can't be changed because another record refers to it", 409)

// 		case "23505":
// 			// unique constraint violation
// 			return CreateCustomError("This record contains duplicated data that conflicts with what is already in the database", 409)

// 		case "23514":
// 			// check constraint violation
// 			return CreateCustomError("This record contains inconsistent or out-of-range data", 400)
// 		}
// 	}
// 	return CreateCustomError("Internal server error", 500)
// }
