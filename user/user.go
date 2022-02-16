package user

type User struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true,unique:true"`
	Username string `json:"username" gorm:"type:varchar(100);unique;not null;default:null"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null;default:null"`
}
