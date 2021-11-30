package user

import (
	"errors"
	"strconv"
)

type User struct {
	Id      int    `json: "id"`
	Name    string `json: "name"`
	Age     int    `json: "age"`
	Friends []int  `json: "friends"`
}

func (u *User) ToString() string {
	return string(u.Id) + " " + u.Name + " " + strconv.Itoa(u.Age) + " [" + strconv.Itoa(len(u.Friends)) + " друга]"
}

func (u *User) AddFriend(userId int) error {
	index := findIndex(u.Friends, userId)
	if index == -1 {
		u.Friends = append(u.Friends, userId)
	} else {
		return errors.New("У пользователя " + u.Name + " уже есть такой друг")
	}
	return nil
}

func (u *User) DeleteFriend(userId int) error {
	if len(u.Friends) < 1 {
		return errors.New("У пользователя нет друзей")
	}
	index := findIndex(u.Friends, userId)
	if index > -1 {
		u.Friends = append(u.Friends[:index], u.Friends[index+1:]...)
	} else {
		return errors.New("У пользователя не найден такой друг")
	}
	return nil
}

func findIndex(list []int, userId int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == userId {
			return i
		}
	}
	return -1
}
