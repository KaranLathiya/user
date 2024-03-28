package dal

import "database/sql"

func Create(db *sql.DB) error { 
	_, err := db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	return nil
}