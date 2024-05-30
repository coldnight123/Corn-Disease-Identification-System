package models

type User struct {
	ID       int64
	UserID   int64 `gorm:"column:user_id"`
	Username string
	Password string
	Email    string `validate:"email"`
	Token    string `gorm:"-"`
}

func (User) TableName() string {
	return "user"
}
