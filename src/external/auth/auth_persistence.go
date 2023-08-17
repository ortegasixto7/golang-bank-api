package auth

import "sync"

type IAuthPersistence interface {
	Create(data *Auth, waitGroup *sync.WaitGroup)
	Update(data *Auth)
	GetByUserNameOrNil(userName string) *Auth
}
