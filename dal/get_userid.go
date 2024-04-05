package dal

import (
	"database/sql"
	"fmt"
	"strings"
	error_handling "user/error"
	"user/model/request"

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
			return "",error_handling.UserDoesNotExist
		}
		return "",error_handling.InternalServerError
	}
	return userID, nil
}
