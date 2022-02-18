package auth

import (
	"errors"
	u "group-project1/entities/user"
	"group-project1/repository/hash"

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
	var user u.Users
	a.db.Model(&user).Where("email = ?", email).First(&user)
	isMatched := hash.CheckPasswordHash(password, user.Password)
	if !isMatched {
		return u.Users{}, errors.New("email atau password tidak valid")
	}
	return user, nil
}

// ===================================================================
