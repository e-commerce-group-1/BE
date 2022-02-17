package auth

import (
	u "group-project1/entities/user"

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
	if err := a.db.Model(&u.Users{}).Where("email = ? AND password = ?", email, password).First(&loggedInUser).Error; err != nil {
		return loggedInUser, err
	}
	return loggedInUser, nil
}

// ===================================================================
