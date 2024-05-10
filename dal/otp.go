package dal

import (
	"database/sql"
	"fmt"
	"strings"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/utils"

	"github.com/jmoiron/sqlx"
)

func StoreOTP(db *sql.DB, storeOTP request.StoreOTP) error {
	var err error
	_, err = db.Exec("INSERT INTO public.otp (email,phone_number,country_code,otp,expires_at,event_type,organization_id) VALUES ( $1 , $2 , $3 , $4 , $5 , $6 , $7 )", storeOTP.Email, storeOTP.PhoneNumber, storeOTP.CountryCode, storeOTP.HashedOTP, utils.AddMinutesToCurrentUTCTime(5), storeOTP.EventType, storeOTP.OrganizationID)
	fmt.Println(err)
	if err != nil {
		return error_handling.InternalServerError
	}
	return nil
}

func VerifyOTP(db *sql.DB, verifyOTP request.VerifyOTP) error {
	var where []string
	var filterArgsList []interface{}
	where = append(where, "event_type = ? ")
	filterArgsList = append(filterArgsList, verifyOTP.EventType)
	if verifyOTP.SignupMode == constant.SIGNUP_MODE_EMAIL || verifyOTP.SignupMode == constant.SIGNUP_MODE_GOOGLE_LOGIN {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.SignupMode == constant.SIGNUP_MODE_PHONE_NUMBER {
		where = append(where, "phone_number = ?", "country_code = ?")
		filterArgsList = append(filterArgsList, verifyOTP.PhoneNumber, verifyOTP.CountryCode)
	}
	if verifyOTP.EventType == constant.EVENT_TYPE_ORGANIZATION_DELETE {
		where = append(where, "organization_id = ?")
		filterArgsList = append(filterArgsList, verifyOTP.OrganizationID)
	}
	query := fmt.Sprintf("SELECT otp, case WHEN current_timestamp() > expires_at THEN true ELSE false END AS otp_expired FROM public.otp WHERE %v", strings.Join(where, " AND "))
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
	where = append(where, "event_type = ? ")
	filterArgsList = append(filterArgsList, verifyOTP.EventType)
	if verifyOTP.SignupMode == constant.SIGNUP_MODE_EMAIL || verifyOTP.SignupMode == constant.SIGNUP_MODE_GOOGLE_LOGIN {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.SignupMode == constant.SIGNUP_MODE_PHONE_NUMBER {
		where = append(where, "phone_number = ?", "country_code = ?")
		filterArgsList = append(filterArgsList, verifyOTP.PhoneNumber, verifyOTP.CountryCode)
	}
	if verifyOTP.EventType == constant.EVENT_TYPE_ORGANIZATION_DELETE {
		where = append(where, "organization_id = ?")
		filterArgsList = append(filterArgsList, verifyOTP.OrganizationID)
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
