package admin

import "gorm.io/gorm"

type Init interface {
	Getone(id uint64) (Admin, error)
	Create(obj Admin) (Admin, error)
	Login(obj Admin) (Admin, error)
}

type adminReciever struct {
	gorm *gorm.DB
}

func NewAdmin(gorm *gorm.DB) Init {
	return adminReciever{
		gorm: gorm,
	}
}
