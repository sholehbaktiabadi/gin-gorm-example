package user

func (r userReciever) GetoneByEmail(email string) (User, error) {
	var user = User{}
	if err := r.gorm.Where("email= ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
