package user

type IUserPersistence interface {
	Create(data *User)
	GetByUserNameOrNil(userName string) *User
}
