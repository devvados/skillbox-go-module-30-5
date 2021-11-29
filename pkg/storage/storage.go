package storage

import (
	"errors"
	"skillbox/module30/skillbox-go-module-30-5/pkg/user"
)

//Тип - хранилище
type Storage map[int]*user.User

//Инициализация хранилища
func (s Storage) NewStorage() Storage {
	return make(map[int]*user.User)
}

//Добавление пользователя в хранилище
func (s Storage) Add(user *user.User) {
	s[user.Id] = user
}

//Получение пользователя из хранилища по имени
func (s Storage) Get(userId int) (*user.User, error) {
	u, ok := s[userId]
	if !ok {
		return nil, errors.New("Пользователь не найден!")
	} else {
		return u, nil
	}
}

//Удаление пользователя из хранилища + у всех из друзей
func (s Storage) Delete(userId int) error {
	_, err := s.Get(userId)
	if err != nil {
		return err
	}

	//Сначала удаление из друзей
	// for _, u := range s.St {
	// 	u.DeleteFriend(userId)
	// }
	//Затем удаление из списка совсем
	s[userId] = nil

	return nil
}
