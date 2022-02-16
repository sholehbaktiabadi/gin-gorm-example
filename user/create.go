package user

func (r userReciever) Create(obj User) (User, error) {
	var user = User{
		Username: obj.Username,
		Name:     obj.Name,
		Email:    obj.Email,
	}
	if err := r.gorm.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
