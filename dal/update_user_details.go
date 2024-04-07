package dal

import (
	"database/sql"
	error_handling "user/error"
	"user/model/request"
	"user/utils"

	"github.com/lib/pq"
)

func UpdateUserPrivacy(db *sql.DB, updateUserPrivacy request.UpdateUserPrivacy, userID string) error {
	_, err := db.Exec("UPDATE public.users SET privacy=$1,updated_at=$2 WHERE id=$3 ;", updateUserPrivacy.Privacy,utils.CurrentUTCTime(0), userID)
	if err != nil {
		return error_handling.InternalServerError
	}
	return nil
}

func UpdateUserNameDetails(db *sql.DB, updateUserNameDetails request.UpdateUserNameDetails, userID string) error {
	_, err := db.Exec("UPDATE public.users SET firstname=$1,lastname=$2,lastname=$3,username=$4,updated_at=$5 WHERE id=$6 ;", updateUserNameDetails.Firstname, updateUserNameDetails.Lastname, updateUserNameDetails.Firstname+" "+updateUserNameDetails.Lastname, updateUserNameDetails.Username,utils.CurrentUTCTime(0), userID)
	if err != nil {
		if dbErr, ok := err.(*pq.Error); ok {
			errCode := dbErr.Code
			switch errCode {
			case "23505":
				// unique constraint violation
				return  error_handling.UsernameAlreadyTaken
			}
			return error_handling.InternalServerError
		}
	}
	return nil
}
