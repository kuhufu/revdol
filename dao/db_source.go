package dao

import (
	"revdol/model"
)

type DBSource struct {
	source Source
}

func (s *DBSource) Register(account *model.Account) (*model.Account, error) {
	return s.source.Register(account)
}

func (s *DBSource) Login(identity, password string) (*model.Account, error) {
	return s.source.Login(identity, password)
}

func (s *DBSource) ChangePassword(id uint, newPwd string) (*model.Account, error) {
	return s.source.ChangePassword(id, newPwd)
}


