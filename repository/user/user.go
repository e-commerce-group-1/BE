package user

import (
	u "group-project1/enitities/user"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// ======================== User Register ==================================
func (ur *UserRepository) Insert(newUser u.User) (u.User, error) {
	if err := ur.db.Save(&newUser).Error; err != nil {
		log.Warn("Found database error:", err)
		return newUser, err
	}

	return newUser, nil
}
// ============================================================================