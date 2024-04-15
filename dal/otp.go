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

func StoreOTP(db *sql.DB, storeOTP request.StoreOTP) error {
	var err error

	if storeOTP.LoginType == "email" || storeOTP.LoginType == "google_login" {
		storeOTP.CountryCode = nil
		storeOTP.PhoneNumber = nil
	} else {
		storeOTP.Email = nil
	}
	_, err = db.Exec("INSERT INTO public.otp (email,phone_number,country_code,otp,expires_at,event_type) VALUES ( $1 , $2 , $3 , $4 , $5 , $6 )", storeOTP.Email, storeOTP.PhoneNumber, storeOTP.CountryCode, storeOTP.HashedOTP, utils.CurrentUTCTime(5), storeOTP.EventType)
	fmt.Println(err)
	if err != nil {
		return error_handling.InternalServerError
	}
	return nil
}

func VerifyOTP(db *sql.DB, verifyOTP request.VerifyOTP) error {
	var where []string
	var filterArgsList []interface{}
	if verifyOTP.SignupMode == "email" || verifyOTP.SignupMode == "google_login" {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.SignupMode == "phone_number" {
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

func DeleteOTPs(db *sql.DB, verifyOTP request.VerifyOTP) error {
	var where []string
	var filterArgsList []interface{}
	if verifyOTP.EventType == "email" || verifyOTP.EventType == "google_login" {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.EventType == "phone_number" {
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
