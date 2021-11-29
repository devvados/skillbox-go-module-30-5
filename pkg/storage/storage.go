package storage

import (
	"errors"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

type Getter interface {
	Get(userId int) *user.User
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
		return nil, errors.New("пользователь не найден")
	} else {
		return u, nil
	}
}

func (r *Repo) Delete(userId int) error {
	_, err := r.Get(userId)
	if err != nil {
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
