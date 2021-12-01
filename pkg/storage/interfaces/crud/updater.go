package crud

type Updater interface {
	UpdateUserAge(userId int, age int) error
	LinkUsers(userLinkFrom int, userLinkTo int) error
}
