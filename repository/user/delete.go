package user

func (r userReciever) Delete(id uint64) (User, error) {
	var user = User{}
	if err := r.gorm.First(&user, id).Error; err != nil {
		return user, err
	}
	if err := r.gorm.Delete(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}
