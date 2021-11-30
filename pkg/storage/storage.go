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
	Add(user *user.User)
}

type Deleter interface {
	Delete(userId int) error
}

type Linker interface {
	Link(userLinkFrom int, userLinkTo int) error
}

type Updater interface {
	UpdateAge(userId int, age int) error
}

type Repo struct {
	Items map[int]*user.User
}

func New() *Repo {
	return &Repo{
		make(map[int]*user.User),
	}
}

func (r *Repo) Add(user *user.User) {
	r.Items[user.Id] = user
}

func (r *Repo) Get(userId int) (*user.User, error) {
	u, ok := r.Items[userId]
	if !ok {
		return nil, errors.New("Пользователь не найден")
	} else {
		return u, nil
	}
}

func (r *Repo) GetAll() []*user.User {
	users := make([]*user.User, 0)
	if len(r.Items) > 0 {
		for _, u := range r.Items {
			users = append(users, u)
		}
	}
	return users
}

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

func (r *Repo) Delete(userId int) error {
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

func (r *Repo) UpdateAge(userId int, age int) error {
	if _, err := r.Get(userId); err != nil {
		return err
	}
	r.Items[userId].Age = age

	return nil
}

func (r *Repo) Link(userLinkFrom int, userLinkTo int) error {
	linkFrom, _ := r.Get(userLinkFrom)
	linkTo, _ := r.Get(userLinkTo)

	if linkFrom == nil || linkTo == nil {
		return errors.New("Один или оба из пользователя не найдены")
	}

	linkFrom.Friends = append(linkFrom.Friends, userLinkTo)
	linkTo.Friends = append(linkTo.Friends, userLinkFrom)
	return nil
}
