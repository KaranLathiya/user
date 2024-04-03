package repository

import (
	"database/sql"
	"user/dal"
	"user/model/request"
)

type Repository interface {
	Signup(signup request.Signup) (string,error)
	StoreOTP(signup request.Signup, otp string, eventType string) error
}

type Repositories struct {
	db *sql.DB
}

// InitRepositories should be called in main.go
func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{db: db}
}

func (r *Repositories) Signup(signup request.Signup) (string,error) {
	return dal.Signup(r.db,signup)
}

func (r *Repositories) StoreOTP(signup request.Signup,otp string, eventType string) error {
	return dal.StoreOTP(r.db,signup,otp,eventType)
}

