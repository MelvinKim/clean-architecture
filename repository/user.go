package repository

import (
	"github.com/clean-architecture/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(name string, email string) (*entity.User, error) {
	user := &entity.User{Name: name, Email: email}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(id int) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ListUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
