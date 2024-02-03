package repositories

import (
	"Hexagonal-Model/models"
	"database/sql"
)

type RepositoryPort interface {
	PotsRegister(register models.RequestRegister) error
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) PotsRegister(register models.RequestRegister) error {

	return nil
}
