package repository

import (
	"errors"
	"go-users/internal/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	db.AutoMigrate(&domain.User{})
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	users := []domain.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id int64) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *domain.User) error {
	res := r.db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(user)
	if res.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return res.Error
}

func (r *userRepository) Delete(id int64) error {
	res := r.db.Delete(&domain.User{}, id)
	if res.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return res.Error
}
