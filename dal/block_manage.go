package dal

import (
	"database/sql"
	"fmt"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"

	"github.com/lib/pq"
)

func BlockUser(db *sql.DB, blockUser request.BlockUser, userID string) error {
	_, err := db.Exec("INSERT INTO public.blocked_user (blocker,blocked,blocked_at) VALUES ($1,$2,$3)", userID, blockUser.BlockedUser, utils.CurrentUTCTime(0))
	fmt.Println(err)
	if err != nil {
		if dbErr, ok := err.(*pq.Error); ok {
			errCode := dbErr.Code
			switch errCode {
			case "23505":
				// unique constraint violation
				return error_handling.AlreadyBlocked
			case "23503":
				// foreign key constraint violation
				return error_handling.InvalidDetails
			}
			return error_handling.InternalServerError
		}
	}
	return nil
}

func UnblockUser(db *sql.DB, blockedUser request.BlockUser, userID string) error {
	result, err := db.Exec("DELETE FROM public.blocked_user WHERE blocker = $1 AND blocked = $2", userID, blockedUser.BlockedUser)
	fmt.Println(err)
	if err != nil {
		return error_handling.InternalServerError
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return error_handling.InternalServerError
	}
	if rowsAffected == 0 {
		return error_handling.InvalidDetails
	}
	return nil
}

func BlockedUserList(db *sql.DB, userID string) ([]response.BlockUserDetails,error) {
	rows, err := db.Query("SELECT blocked,blocked_at FROM public.blocked_user WHERE blocker = $1 ", userID)
	if err != nil {
		return nil,error_handling.InternalServerError
	}
	var blockedUserList []response.BlockUserDetails
	for rows.Next(){
		var blockedUser response.BlockUserDetails
		err := rows.Scan(&blockedUser.BlockedUser,&blockedUser.BlockedAt)
		if err != nil {
			return nil,error_handling.InternalServerError
		}
		blockedUserList = append(blockedUserList, blockedUser)
	}
	defer rows.Close()
	return blockedUserList,nil
}