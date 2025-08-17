package domain

type User struct {
	 ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	 Name string `gorm:"size:255;not null" json:"name"`
	 Email string `gorm:"size:255;unique;not null" json:"email"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int64) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

type UserService interface {
	GetAllUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int64) error
}