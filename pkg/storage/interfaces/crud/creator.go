package crud

import "skillbox/module30/skillbox-go-module-30-5/pkg/model"

type Creator interface {
	AddUser(user *model.User)
}
