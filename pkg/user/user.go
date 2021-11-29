package user

import (
	"errors"
	"strconv"
)

//Пользователь
type User struct {
	Id      int    `'json:"id`
	Name    string `json: "name"`
	Age     int    `json: "age"`
	Friends []int  `json: "friends"`
}

//Инфо о пользователе в виде строки
func (u *User) ToString() string {
	return u.Name + " " + strconv.Itoa(u.Age) + " [" + strconv.Itoa(len(u.Friends)) + " друга]"
}

//Добавление друга пользователю
func (u *User) AddFriend(userId int) error {
	index := findIndex(u.Friends, userId)
	if index == -1 {
		u.Friends = append(u.Friends, userId)
	} else {
		return errors.New("У пользователя " + u.Name + " уже есть такой друг")
	}
	return nil
}

//Удаление друга у пользователя
func (u *User) DeleteFriend(userId int) {
	if len(u.Friends) < 1 {
		return
	}

	index := findIndex(u.Friends, userId)
	if index > -1 {
		u.Friends = append(u.Friends[:index], u.Friends[index+1:]...)
	}
}

//Поиск индекса в списке пользователей для удаления из друзей
func findIndex(list []int, userId int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == userId {
			return i
		}
	}
	return -1
}
