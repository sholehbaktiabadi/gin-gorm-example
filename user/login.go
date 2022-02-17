package user

func (r userReciever) Login(obj User) (User, error) {
	var user = User{}
	if err := r.gorm.Where(&User{Email: obj.Email}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
