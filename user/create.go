package user

func (r userReciever) Create(obj User) (User, error) {
	if err := r.gorm.Create(&obj).Error; err != nil {
		return obj, err
	}
	return obj, nil
}
