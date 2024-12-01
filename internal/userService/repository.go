package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(email, password User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, email User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) CreateUser(email User, password User) (User, error) {
	user := User{Email: email.Email, Password: password.Password}
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r userRepository) DeleteUserByID(id uint) error {
	err := r.db.Delete(&User{}, id).Error
	return err
}

func (r userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r userRepository) UpdateUserByID(id uint, email User) (User, error) {
	err := r.db.Model(&email).Where("id = ?", id).Updates(email).Error
	return email, err
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
