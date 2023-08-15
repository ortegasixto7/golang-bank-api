package auth

type Auth struct {
	Id        string
	UserName  string
	Password  string
	CreatedAt int64
}

func (Auth) TableName() string {
	return "auth"
}
