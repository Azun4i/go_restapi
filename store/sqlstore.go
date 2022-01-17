package store

import (
	"database/sql"
	"errors"
)

type Store struct {
	db       *sql.DB
	userRepo *UserRepository
}

var (
	err = errors.New("not authenticated")
)

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() *UserRepository {

	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepository{
		store: s,
	}

	return s.userRepo
}
