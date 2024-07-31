package dbusers

import "gorm.io/gorm"

type UserRepositoryAdapter struct {
	DB *gorm.DB
}

func (adapter UserRepositoryAdapter) Save(model UserModel) error {
	result := adapter.DB.Create(model)
	return result.Error
}
