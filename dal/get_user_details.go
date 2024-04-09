package dal

import (
	"database/sql"
	"fmt"
	"strings"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"

	"github.com/jmoiron/sqlx"
)

func GetUserID(db *sql.DB, verifyOTP request.VerifyOTP) (string, error) {
	var where []string
	var filterArgsList []interface{}
	var userID string
	if verifyOTP.SignupMode == "email" || verifyOTP.SignupMode == "google_login" {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.SignupMode == "phone" {
		where = append(where, "phone_number = ?", "country_code = ?")
		filterArgsList = append(filterArgsList, verifyOTP.PhoneNumber, verifyOTP.CountryCode)
	}
	query := fmt.Sprintf("SELECT id FROM public.users WHERE %v", strings.Join(where, " AND "))
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	// fmt.Println(query)
	err := db.QueryRow(query, filterArgsList...).Scan(&userID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", error_handling.UserDoesNotExist
		}
		return "", error_handling.InternalServerError
	}
	return userID, nil
}

func GetUsernameByID(db *sql.DB, id string) (string, error) {
	var username string
	err := db.QueryRow("SELECT username FROM public.users WHERE id=$1", id).Scan(&username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", error_handling.UserDoesNotExist
		}
		return "", error_handling.InternalServerError
	}
	return username, nil
}

func GetUserDetailsByID(db *sql.DB, id string) (response.UserDetails, error) {
	userDetails:=response.UserDetails{
		UserID: id,
	}
	err := db.QueryRow("SELECT firstname,lastname,fullname,username,email,phone_number,country_code,privacy,created_at,updated_at FROM public.users WHERE id=$1", id).Scan(&userDetails.Firstname,&userDetails.Lastname,&userDetails.Fullname,&userDetails.Username,&userDetails.Email,&userDetails.PhoneNumber,&userDetails.CountryCode,&userDetails.Privacy,&userDetails.CreatedAt,&userDetails.UpdatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return userDetails, error_handling.UserDoesNotExist
		}
		return userDetails, error_handling.InternalServerError
	}
	return userDetails, nil
}
