package dal

import "database/sql"

func Delete(db *sql.DB) error { 
	_, err := db.Exec("DELETE users where id = $1")
	if err != nil {
		return err
	}
	return nil
}