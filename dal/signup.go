package dal

import (
	"database/sql"
	"user/model/request"
)

func Signup(db *sql.DB, signup request.Signup) (string,error) {
	// db.QueryRow()
	return "",nil
}
