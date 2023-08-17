package user

import "sync"

type IUserPersistence interface {
	Create(data *User, waitGroup *sync.WaitGroup)
	GetByUserNameOrNil(userName string) *User
}
