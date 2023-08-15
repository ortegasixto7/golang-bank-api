package postgres

import (
	"fmt"

	"github.com/ortegasixto7/go-bank-api/src/core/user"
)

type PostgresUserPersistence struct{}

func (this PostgresUserPersistence) Create(data *user.User) {

	fmt.Printf("%v \n", Database)

	result := Database.Create(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("User Created!")
	}
}

func (this PostgresUserPersistence) GetByUserNameOrNil(userName string) *user.User {
	return nil
}
