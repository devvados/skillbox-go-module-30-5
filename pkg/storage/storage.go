package storage

import (
	"errors"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

type Getter interface {
	Get(userId int) (*user.User, error)
	GetAll() []*user.User
	GetFriends(userId int) ([]*user.User, error)
}

type Adder interface {
	AddUser(user *user.User)
}

type Deleter interface {
	DeleteUser(userId int) error
}

type Linker interface {
	LinkUsers(userLinkFrom int, userLinkTo int) error
}

type Updater interface {
	UpdateUserAge(userId int, age int) error
}

type Repo struct {
	Items map[int]*user.User
}

func New() *Repo {
	return &Repo{
		make(map[int]*user.User),
	}
}

//Добавление пользователя в хранилище
func (r *Repo) AddUser(user *user.User) {
	r.Items[user.Id] = user
}

//Получение пользователя по идентификатору
func (r *Repo) Get(userId int) (*user.User, error) {
	u, ok := r.Items[userId]
	if !ok {
		return nil, errors.New("Пользователь не найден")
	} else {
		return u, nil
	}
}

//Получение всех пользователей в хранилище
func (r *Repo) GetAll() []*user.User {
	users := make([]*user.User, 0)
	if len(r.Items) > 0 {
		for _, u := range r.Items {
			users = append(users, u)
		}
	}
	return users
}

//Получение друзей пользователя
func (r *Repo) GetFriends(userId int) ([]*user.User, error) {
	u, err := r.Get(userId)
	if err != nil {
		return nil, err
	}
	friends := make([]*user.User, 0)
	for _, val := range u.Friends {
		friends = append(friends, r.Items[val])
	}
	return friends, nil
}

//Удаление пользователя из хранилища
func (r *Repo) DeleteUser(userId int) error {
	if _, err := r.Get(userId); err != nil {
		return err
	}
	//Сначала удаление из друзей
	for _, u := range r.Items {
		u.DeleteFriend(userId)
	}
	//Затем удаление из списка совсем
	delete(r.Items, userId)

	return nil
}

//Обновление возраста пользователя
func (r *Repo) UpdateUserAge(userId int, age int) error {
	if _, err := r.Get(userId); err != nil {
		return err
	}
	r.Items[userId].Age = age

	return nil
}

//Добавление пользователя в друзья
func (r *Repo) LinkUsers(userLinkFrom int, userLinkTo int) error {
	linkFrom, _ := r.Get(userLinkFrom)
	linkTo, _ := r.Get(userLinkTo)

	if linkFrom == nil || linkTo == nil {
		return errors.New("Один или оба из пользователя не найдены")
	}

	linkFrom.Friends = append(linkFrom.Friends, userLinkTo)
	linkTo.Friends = append(linkTo.Friends, userLinkFrom)
	return nil
}
