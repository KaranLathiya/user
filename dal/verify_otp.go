package dal

import (
	"database/sql"
	"fmt"
	"strings"
	error_handling "user/error"
	"user/model/request"
	"user/utils"

	"github.com/jmoiron/sqlx"
)

func VerifyOTP(db *sql.DB, verifyOTP request.VerifyOTP) error {
	var where []string
	var filterArgsList []interface{}

	if verifyOTP.SignupMode == "email" || verifyOTP.SignupMode == "google_login" {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.SignupMode == "phone" {
		where = append(where, "phone_number = ?", "country_code = ?")
		filterArgsList = append(filterArgsList, verifyOTP.PhoneNumber, verifyOTP.CountryCode)
	}
	query := fmt.Sprintf("SELECT otp, case WHEN '%s' > expires_at THEN true ELSE false END AS otp_expired FROM public.otp WHERE %v", utils.CurrentUTCTime(0), strings.Join(where, " AND "))
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	fmt.Println(query)
	rows, err := db.Query(query, filterArgsList...)
	if err != nil {
		return error_handling.InternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var storedOTP string
		var isOTPExpired bool
		err := rows.Scan(&storedOTP, &isOTPExpired)
		if err != nil {
			return error_handling.InternalServerError
		}
		success := utils.CompareHashAndPassword([]byte(storedOTP), []byte(verifyOTP.OTP))
		if success {
			if isOTPExpired {
				return error_handling.ExpiredOTP
			}
			return nil
		}
	}
	return error_handling.InvalidOTP
}
