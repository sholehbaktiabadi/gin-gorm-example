package admin

func (r adminReciever) Getone(id uint64) (Admin, error) {
	var admin = Admin{}
	if err := r.gorm.First(&admin, id).Error; err != nil {
		return admin, err
	}
	return admin, nil
}
