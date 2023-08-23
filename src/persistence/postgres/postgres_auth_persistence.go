package postgres

import (
	"fmt"
	"sync"

	"github.com/ortegasixto7/go-bank-api/src/external/auth"
)

type PostgresAuthPersistence struct{}

func (this *PostgresAuthPersistence) Create(data *auth.Auth, waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()
	}

	result := Database.Create(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("Auth Created!")
	}
}

func (this *PostgresAuthPersistence) Update(data *auth.Auth) {
	result := Database.Create(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func (this *PostgresAuthPersistence) GetByUserNameOrNil(userName string) *auth.Auth {
	var result *auth.Auth
	Database.Where(&auth.Auth{UserName: userName}).First(&result)
	if result.Id == "" {
		return nil
	}
	return result
}
