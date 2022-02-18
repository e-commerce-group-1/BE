package admin

import (
	"group-project1/entities/user"
	"group-project1/repository/hash"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

// ======================== Admin Register ==================================
func (ur *AdminRepository) Insert(NewAdmin user.Users) (user.Users, error) {
	NewAdmin.Password, _ = hash.HashPassword(NewAdmin.Password)
	if err := ur.db.Create(&NewAdmin).Error; err != nil {
		return NewAdmin, err
	}
	return NewAdmin, nil
}
