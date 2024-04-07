package repository

import (
	"database/sql"
	"user/dal"
	"user/model/request"
	"user/model/response"
)

type Repository interface {
	StoreOTP(storeOTP request.StoreOTP) error
	VerifyOTP(verifyOTP request.VerifyOTP) error
	DeleteOTPs(verifyOTP request.VerifyOTP) error
	UserCreate(verifyOTP request.VerifyOTP) (string, error)
	GetUserID(verifyOTP request.VerifyOTP) (string, error)
	UpdateUserPrivacy(updateUserPrivacy request.UpdateUserPrivacy, userID string) error
	UpdateUserNameDetails(updateUserNameDetails request.UpdateUserNameDetails, userID string) error
	BlockUser(blockUser request.BlockUser, userID string) error
	UnblockUser(blockedUser request.BlockUser, userID string) error
	BlockedUserList(userID string) ([]response.BlockUserDetails, error)
}

type Repositories struct {
	db *sql.DB
}

// InitRepositories should be called in main.go
func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{db: db}
}

func (r *Repositories) StoreOTP(storeOTP request.StoreOTP) error {
	return dal.StoreOTP(r.db, storeOTP)
}

func (r *Repositories) VerifyOTP(verifyOTP request.VerifyOTP) error {
	return dal.VerifyOTP(r.db, verifyOTP)
}

func (r *Repositories) DeleteOTPs(verifyOTP request.VerifyOTP) error {
	return dal.DeleteOTPs(r.db, verifyOTP)
}

func (r *Repositories) UserCreate(verifyOTP request.VerifyOTP) (string, error) {
	return dal.UserCreate(r.db, verifyOTP)
}

func (r *Repositories) GetUserID(verifyOTP request.VerifyOTP) (string, error) {
	return dal.GetUserID(r.db, verifyOTP)
}

func (r *Repositories) UpdateUserPrivacy(updateUserPrivacy request.UpdateUserPrivacy, userID string) error {
	return dal.UpdateUserPrivacy(r.db, updateUserPrivacy, userID)
}

func (r *Repositories) UpdateUserNameDetails(updateUserNameDetails request.UpdateUserNameDetails, userID string) error {
	return dal.UpdateUserNameDetails(r.db, updateUserNameDetails, userID)
}

func (r *Repositories) BlockUser(blockUser request.BlockUser, userID string) error {
	return dal.BlockUser(r.db, blockUser, userID)
}

func (r *Repositories) UnblockUser(blockedUser request.BlockUser, userID string) error {
	return dal.UnblockUser(r.db, blockedUser, userID)
}

func (r *Repositories) BlockedUserList(userID string) ([]response.BlockUserDetails, error) {
	return dal.BlockedUserList(r.db, userID)
}
