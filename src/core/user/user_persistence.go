package user

import "sync"

type UserPersistence interface {
	Create(data *User, waitGroup *sync.WaitGroup)
	GetByUserNameOrNil(userName string) *User
}
