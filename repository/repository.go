package repository

import (
	"database/sql"
	"user/dal"
	"user/model/request"
)

type Repository interface {
	StoreOTP(signup request.Signup, otp string, eventType string) error
	VerifyOTP(verifyOTP request.VerifyOTP) error 
	DeleteOTPs(verifyOTP request.VerifyOTP) error 
	UserCreate(verifyOTP request.VerifyOTP) (string,error)
	GetUserID(verifyOTP request.VerifyOTP) (string,error)
}

type Repositories struct {
	db *sql.DB
}

// InitRepositories should be called in main.go
func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{db: db}
}

func (r *Repositories) StoreOTP(signup request.Signup,otp string, eventType string) error {
	return dal.StoreOTP(r.db,signup,otp,eventType)
}

func (r *Repositories) VerifyOTP(verifyOTP request.VerifyOTP) error {
	return dal.VerifyOTP(r.db,verifyOTP)
}

func (r *Repositories) DeleteOTPs(verifyOTP request.VerifyOTP) error {
	return dal.DeleteOTPs(r.db,verifyOTP)
}

func (r *Repositories) UserCreate(verifyOTP request.VerifyOTP) (string,error) {
	return dal.UserCreate(r.db,verifyOTP)
}

func (r *Repositories) GetUserID(verifyOTP request.VerifyOTP) (string,error) {
	return dal.GetUserID(r.db,verifyOTP)
}
