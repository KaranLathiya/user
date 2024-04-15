package dal

import (
	"database/sql"
	"fmt"
	"strings"
	"user/model/request"
	"user/model/response"

	"github.com/jmoiron/sqlx"
)

func GetUserList(db *sql.DB, userID string, userListParameter request.UserListParameter) ([]response.User, error) {
	var where []string
	var filterArgsList []interface{}

	if *userListParameter.Fullname != "" {
		where = append(where, "fullname ILIKE '%' || ? || '%'")
		filterArgsList = append(filterArgsList, *userListParameter.Fullname)
	} else if *userListParameter.Email != "" {
		where = append(where, "email ILIKE '%' || ? || '%'")
		filterArgsList = append(filterArgsList, *userListParameter.Email)
	} else if *userListParameter.PhoneNumber != "" {
		where = append(where, "phone_number ILIKE '%' || ? || '%'")
		filterArgsList = append(filterArgsList, *userListParameter.PhoneNumber)
	}

	if userListParameter.OrderBy == "" {
		userListParameter.OrderBy = "fullname"
	} else if userListParameter.OrderBy == "date" {
		userListParameter.OrderBy = "created_at"
	}

	if userListParameter.OrderBy == "" {
		userListParameter.OrderBy = "asc"
	}

	joinCondition := "LEFT JOIN blocked_user b1 ON u.id = b1.blocked AND b1.blocker = '" + userID + "' LEFT JOIN blocked_user b2 ON u.id = b2.blocker AND b2.blocked = '" + userID + "' WHERE b1.blocker IS NULL AND b2.blocked IS NULL AND u.id != '" + userID + "'"
	var andKeyword string
	if len(where) > 0 {
		andKeyword = "AND"
	}

	query := fmt.Sprintf("SELECT u.id, firstname, lastname, fullname, username, email, phone_number, country_code FROM public.users u %s AND privacy = 'public' %s %v ORDER BY %s %s LIMIT %d OFFSET %d", joinCondition, andKeyword, strings.Join(where, " AND "), userListParameter.OrderBy, userListParameter.Order, userListParameter.Limit, userListParameter.Offset)
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
