package admin

func (r adminReciever) Login(obj Admin) (Admin, error) {
	var admin Admin
	if err := r.gorm.Where(&Admin{Email: obj.Email}).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}
