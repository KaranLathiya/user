package repository

import (
	"database/sql"
	"go-structure/dal"
)

type Repository interface {
	Create() error
	Delete() error
}

type Repositories struct {
	db *sql.DB
}

// InitRepositories should be called in main.go
func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{db: db}
}

func (r *Repositories) Create() error {
	return dal.Create(r.db)
}

func (r *Repositories) Delete() error {
	return dal.Delete(r.db)
}
