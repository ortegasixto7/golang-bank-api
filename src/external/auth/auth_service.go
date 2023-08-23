package auth

import "sync"

type AuthService struct {
	AuthPersistence IAuthPersistence
}

func (this *AuthService) Create(data *Auth, waitGroup *sync.WaitGroup) {
	this.AuthPersistence.Create(data, waitGroup)
}

func (this *AuthService) Update(data *Auth) {
	this.AuthPersistence.Update(data)
}

func (this *AuthService) GetByUserNameOrNil(userName string) *Auth {
	return this.AuthPersistence.GetByUserNameOrNil(userName)
}
