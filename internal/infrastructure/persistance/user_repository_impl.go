package persistance

import (
	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	// "github.com/AAteddy/event-booking-api/internal/domain/repositories"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{ DB *gorm.DB }

func (repo *UserRepositoryImpl) Save(user *entities.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepositoryImpl) FindByID(id string) (*entities.User, error) {
	var user entities.User
	err := repo.DB.First(&user, "id = ?", id).Error
	return &user, err
}

func (repo *UserRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (repo *UserRepositoryImpl) FindAll(offset, limit int) ([]*entities.User, error) {
	var users []*entities.User
	err := repo.DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (repo *UserRepositoryImpl) Update(user *entities.User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepositoryImpl) Delete(id string) error {
	return repo.DB.Where("id = ?", id).Delete(&entities.User{}).Error
}

func (repo *UserRepositoryImpl) Count() (int64, error) {
	var count int64
	err := repo.DB.Model(&entities.User{}).Count(&count).Error
	return count, err
}
