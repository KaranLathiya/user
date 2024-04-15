package dal

import (
	"database/sql"
	"fmt"
	"strings"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func GetUserID(db *sql.DB, verifyOTP request.VerifyOTP) (string, error) {
	var where []string
	var filterArgsList []interface{}
	var userID string
	if verifyOTP.SignupMode == "email" || verifyOTP.SignupMode == "google_login" {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, verifyOTP.Email)
	} else if verifyOTP.SignupMode == "phone_number" {
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

func GetIDByUsername(db *sql.DB, username string) (string, error) {
	var id string
	err := db.QueryRow("SELECT id FROM public.users WHERE username = $1", username).Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", error_handling.UserDoesNotExist
		}
		return "", error_handling.InternalServerError
	}
	return id, nil
}

func GetUserDetailsByID(db *sql.DB, id string, userID string) (response.UserDetails, error) {
	userDetails := response.UserDetails{
		UserID: id,
	}
	err := db.QueryRow("SELECT firstname,lastname,fullname,username,email,phone_number,country_code,privacy,created_at,updated_at FROM public.users u LEFT JOIN blocked_user b1 ON u.id = b1.blocked AND b1.blocker = $1 LEFT JOIN blocked_user b2 ON u.id = b2.blocker AND b2.blocked = $1 WHERE b1.blocker IS NULL AND b2.blocked IS NULL AND u.id = $2 AND privacy = 'public' ", userID, id).Scan(&userDetails.Firstname, &userDetails.Lastname, &userDetails.Fullname, &userDetails.Username, &userDetails.Email, &userDetails.PhoneNumber, &userDetails.CountryCode, &userDetails.Privacy, &userDetails.CreatedAt, &userDetails.UpdatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return response.UserDetails{}, nil
		}
		return userDetails, error_handling.InternalServerError
	}
	return userDetails, nil
}

func GetCurrentUserDetailsByID(db *sql.DB, userID string) (response.UserDetails, error) {
	userDetails := response.UserDetails{
		UserID: userID,
	}
	err := db.QueryRow("SELECT firstname,lastname,fullname,username,email,phone_number,country_code,privacy,created_at,updated_at FROM public.users WHERE id = $1 ;", userID).Scan(&userDetails.Firstname, &userDetails.Lastname, &userDetails.Fullname, &userDetails.Username, &userDetails.Email, &userDetails.PhoneNumber, &userDetails.CountryCode, &userDetails.Privacy, &userDetails.CreatedAt, &userDetails.UpdatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return userDetails, error_handling.UserDoesNotExist
		}
		return userDetails, error_handling.InternalServerError
	}
	return userDetails, nil
}

func UpdateUserPrivacy(db *sql.DB, updateUserPrivacy request.UpdateUserPrivacy, userID string) error {
	result, err := db.Exec("UPDATE public.users SET privacy= $1 ,updated_at = $2 WHERE id = $3 ;", updateUserPrivacy.Privacy, utils.CurrentUTCTime(0), userID)
	if err != nil {
		return error_handling.InternalServerError
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return error_handling.InternalServerError
	}
	if rowsAffected == 0 {
		return error_handling.UserDoesNotExist
	}
	return nil
}

func UpdateBasicDetails(db *sql.DB, updateUserNameDetails request.UpdateUserNameDetails, userID string) error {
	_, err := db.Exec("UPDATE public.users SET firstname = $1,lastname = $2, fullname = $3, username = $4, updated_at = $5 WHERE id = $6;", updateUserNameDetails.Firstname, updateUserNameDetails.Lastname, updateUserNameDetails.Firstname+" "+updateUserNameDetails.Lastname, updateUserNameDetails.Username, utils.CurrentUTCTime(0), userID)
	if err != nil {
		if dbErr, ok := err.(*pq.Error); ok {
			errCode := dbErr.Code
			switch errCode {
			case "23505":
				// unique constraint violation
				return error_handling.UsernameAlreadyTaken
			}
			return error_handling.InternalServerError
		}
	}
	return nil
}
