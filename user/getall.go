package user

func (r userReciever) Getall() ([]User, error) {
	var user = []User{}
	if err := r.gorm.Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
