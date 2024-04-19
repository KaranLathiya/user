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
	CreateUser(verifyOTP request.VerifyOTP) (string, error)
	GetUserID(verifyOTP request.VerifyOTP) (string, error)
	UpdateUserPrivacy(updateUserPrivacy request.UpdateUserPrivacy, userID string) error
	UpdateBasicDetails(updateUserNameDetails request.UpdateUserNameDetails, userID string) error
	BlockUser(blockUser request.BlockUser, userID string) error
	UnblockUser(blockedUser request.BlockUser, userID string) error
	BlockedUserList(userID string) ([]response.BlockUserDetails, error)
	GetIDByUsername(username string) (string, error)
	GetCurrentUserDetailsByID(userID string) (response.UserDetails, error)
	GetUserDetailsByID(id string, userID string) (response.UserDetails, error)
	GetUserList(userID string, userListParameter request.UserListParameter) ([]response.User, error)
	UserExistence(email *string, phoneNumber *string) (bool, error)
	GetUsersDetailsByIDs(userIDs []string) (map[string]response.UserDetails, error)
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

func (r *Repositories) CreateUser(verifyOTP request.VerifyOTP) (string, error) {
	return dal.CreateUser(r.db, verifyOTP)
}

func (r *Repositories) GetUserID(verifyOTP request.VerifyOTP) (string, error) {
	return dal.GetUserID(r.db, verifyOTP)
}

func (r *Repositories) UpdateUserPrivacy(updateUserPrivacy request.UpdateUserPrivacy, userID string) error {
	return dal.UpdateUserPrivacy(r.db, updateUserPrivacy, userID)
}

func (r *Repositories) UpdateBasicDetails(updateUserNameDetails request.UpdateUserNameDetails, userID string) error {
	return dal.UpdateBasicDetails(r.db, updateUserNameDetails, userID)
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

func (r *Repositories) GetIDByUsername(username string) (string, error) {
	return dal.GetIDByUsername(r.db, username)
}

func (r *Repositories) GetUserDetailsByID(id string, userID string) (response.UserDetails, error) {
	return dal.GetUserDetailsByID(r.db, id, userID)
}

func (r *Repositories) GetCurrentUserDetailsByID(userID string) (response.UserDetails, error) {
	return dal.GetCurrentUserDetailsByID(r.db, userID)
}

func (r *Repositories) GetUserList(userID string, userListParameter request.UserListParameter) ([]response.User, error) {
	return dal.GetUserList(r.db, userID, userListParameter)
}

func (r *Repositories) UserExistence(email *string, phoneNumber *string) (bool, error) {
	return dal.UserExistence(r.db, email, phoneNumber)
}

func (r *Repositories) GetUsersDetailsByIDs(userIDs []string) (map[string]response.UserDetails, error) {
	return dal.GetUsersDetailsByIDs(r.db, userIDs)
}
