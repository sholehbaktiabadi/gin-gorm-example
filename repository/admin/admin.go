package admin

type Admin struct {
	ID       uint64 `json:"id" gorm:"primaryKey;autoIncrement:true;unique:true"`
	Email    string `json:"email" gorm:"type:varchar(100);unique:true;not null;default:null"`
	Password string `json:"password" gorm:"type:varchar(100);not null;default:null"`
}
