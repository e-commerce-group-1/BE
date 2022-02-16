package auth

import (
	u "group-project1/enitities/user"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// =========================== Login =================================
func (a *AuthRepository) Login(email, password string) (u.Users, error) {
	loggedInUser := u.Users{Email: email, Password: password}

	if err := a.db.First(&loggedInUser).Error; err != nil {
		return loggedInUser, err
	}
	return loggedInUser, nil
}

// ===================================================================
