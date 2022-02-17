package user

import "gorm.io/gorm"

type Init interface {
	Create(obj User) (User, error)
	Login(obj User) (User, error)
	GetoneByEmail(email string) (User, error)
	Getone(id uint64) (User, error)
	Update(obj User, id uint64) (User, error)
	Delete(id uint64) (User, error)
	Getall() ([]User, error)
}

type userReciever struct {
	gorm *gorm.DB
}

func NewUser(gorm *gorm.DB) Init {
	return userReciever{
		gorm: gorm,
	}
}
