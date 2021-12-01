package repositories

import (
	"errors"
	"skillbox/module30/skillbox-go-module-30-5/pkg/model"
	"skillbox/module30/skillbox-go-module-30-5/pkg/storage/interfaces"
)

type LocalRepository struct {
	Items map[int]*model.User
}

func NewLocalRepository() interfaces.Repository {
	return LocalRepository{
		Items: make(map[int]*model.User),
	}
}

//Добавление пользователя в хранилище
func (r LocalRepository) AddUser(user *model.User) {
	r.Items[user.Id] = user
}

//Получение пользователя по идентификатору
func (r LocalRepository) Get(userId int) (*model.User, error) {
	u, ok := r.Items[userId]
	if !ok {
		return nil, errors.New("Пользователь не найден")
	} else {
		return u, nil
	}
}

//Получение всех пользователей в хранилище
func (r LocalRepository) GetAll() []*model.User {
	users := make([]*model.User, 0)
	if len(r.Items) > 0 {
		for _, u := range r.Items {
			users = append(users, u)
		}
	}
	return users
}

//Получение друзей пользователя
func (r LocalRepository) GetFriends(userId int) ([]*model.User, error) {
	u, err := r.Get(userId)
	if err != nil {
		return nil, err
	}
	friends := make([]*model.User, 0)
	for _, val := range u.Friends {
		friends = append(friends, r.Items[val])
	}
	return friends, nil
}

//Удаление пользователя из хранилища
func (r LocalRepository) DeleteUser(userId int) error {
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
func (r LocalRepository) UpdateUserAge(userId int, age int) error {
	if _, err := r.Get(userId); err != nil {
		return err
	}
	r.Items[userId].Age = age

	return nil
}

//Добавление пользователя в друзья
func (r LocalRepository) LinkUsers(userLinkFrom int, userLinkTo int) error {
	linkFrom, _ := r.Get(userLinkFrom)
	linkTo, _ := r.Get(userLinkTo)

	if linkFrom == nil || linkTo == nil {
		return errors.New("Один или оба из пользователя не найдены")
	}

	linkFrom.Friends = append(linkFrom.Friends, userLinkTo)
	linkTo.Friends = append(linkTo.Friends, userLinkFrom)
	return nil
}
