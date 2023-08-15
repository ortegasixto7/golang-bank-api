package user

type User struct {
	Id        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	UserName  string
	CreatedAt int64 `gorm:"autoCreateTime:false"`
}
