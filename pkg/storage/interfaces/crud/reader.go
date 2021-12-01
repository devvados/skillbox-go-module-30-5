package crud

import "skillbox/module30/skillbox-go-module-30-5/pkg/model"

type Reader interface {
	Get(userId int) (*model.User, error)
	GetAll() []*model.User
	GetFriends(userId int) ([]*model.User, error)
}
