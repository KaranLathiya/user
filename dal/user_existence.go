package dal

import (
	"database/sql"
	"fmt"
	"strings"
	error_handling "user/error"

	"github.com/jmoiron/sqlx"
)

func UserExistence(db *sql.DB, email string, phoneNumber string) (bool, error) {
	var where []string
	var filterArgsList []interface{}
	if phoneNumber == "" {
		where = append(where, "email = ? ")
		filterArgsList = append(filterArgsList, email)
	} else if email == "" {
		where = append(where, "phone_number = ?")
		filterArgsList = append(filterArgsList, phoneNumber)
	}
	query := fmt.Sprintf("SELECT id from public.otp WHERE %v", strings.Join(where, " AND "))
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	// fmt.Println(query)
	var id string
	err := db.QueryRow(query, filterArgsList...).Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, error_handling.InternalServerError
	}
	return true, nil
}