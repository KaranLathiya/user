package dal

import (
	"database/sql"
	"fmt"
	"strings"
	"user/model/response"

	"github.com/jmoiron/sqlx"
)

func GetUserList(db *sql.DB, userID string, where []string, filterArgsList []interface{}, orderBy string, order string, limit int, offset int, blockedUserIDs []string) ([]response.User, error) {
	var andKeyword string
	if len(where) > 0 {
		andKeyword = "AND"
	}
	
	query := fmt.Sprintf("SELECT id, firstname, lastname, fullname, username, email, phone_number, country_code FROM public.users WHERE privacy = 'public' %s %v ORDER BY %s %s LIMIT %d OFFSET %d", andKeyword, strings.Join(where, " AND "), orderBy, order, limit, offset)
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := db.Query(query, filterArgsList...)
	if err != nil {
		return nil, err
	}
	var userList []response.User
	for rows.Next() {
		var user response.User
		err = rows.Scan(&user.UserID, &user.Firstname, &user.Lastname, &user.Fullname, &user.Username, &user.Email, &user.PhoneNumber, &user.CountryCode)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}
