package crud

type Deleter interface {
	DeleteUser(userId int) error
}
