package dal

import (
	"database/sql"
	"fmt"
	"strings"
	error_handling "user/error"
	"user/model/request"

	"github.com/jmoiron/sqlx"
)

func DeleteOTPs(db *sql.DB, verifyOTP request.VerifyOTP) error {
	var where []string
	var filterArgsList []interface{}
	if verifyOTP.EventType == "email" || verifyOTP.EventType == "google_login" {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.EventType == "phone" {
		where = append(where, "phone_number = ?", "country_code = ?")
		filterArgsList = append(filterArgsList, verifyOTP.PhoneNumber, verifyOTP.CountryCode)
	}
	query := fmt.Sprintf("DELETE FROM public.otp WHERE %v", strings.Join(where, " AND "))
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	// fmt.Println(query)
	_, err := db.Exec(query, filterArgsList...)
	if err != nil {
		return error_handling.InternalServerError
	}
	return nil
}
