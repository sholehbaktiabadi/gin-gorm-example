package admin

func (r adminReciever) Create(obj Admin) (Admin, error) {
	if err := r.gorm.Create(&obj).Error; err != nil {
		return obj, err
	}
	return obj, nil
}
