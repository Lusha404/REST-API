package teststore

import (
	"HTTP-REST-API/internal/app/model"
	"HTTP-REST-API/internal/app/store"
)

type Store struct {
	UserRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
	return s.UserRepository
}
