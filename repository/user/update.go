package user

func (r userReciever) Update(obj User, id uint64) (User, error) {
	var user = User{}
	if err := r.gorm.First(&user, id).Error; err != nil {
		return user, err
	}
	if err := r.gorm.Model(&user).Updates(obj).Error; err != nil {
		return user, err
	}
	return user, nil
}
