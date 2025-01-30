package repository

import (
	"github.com/richieieie/event-booking/internal/model"
	"gorm.io/gorm"
)

// With bigger projects, we can create a generic repository, they will perform basic function such as get all, get by id, create, update, delete, etc.

type IUserRepository interface {
	CreateOne(user model.User) (int64, error)
	GetByEmailAndPassword(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r userRepository) GetByEmailAndPassword(email string) (model.User, error) {
	var user model.User
	err := r.db.Select("id", "password").Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r userRepository) CreateOne(user model.User) (int64, error) {
	err := r.db.Where("email = ?", user.Email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = r.db.Create(&user).Error
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	return user.Id, nil
}
