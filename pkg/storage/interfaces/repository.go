package interfaces

import "skillbox/module30/skillbox-go-module-30-5/pkg/model"

type Repository interface {
	AddUser(user *model.User)
	DeleteUser(userId int) error
	Get(userId int) (*model.User, error)
	GetAll() []*model.User
	GetFriends(userId int) ([]*model.User, error)
	UpdateUserAge(userId int, age int) error
	LinkUsers(userLinkFrom int, userLinkTo int) error
}
