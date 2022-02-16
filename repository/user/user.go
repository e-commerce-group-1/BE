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

// ======================== Get Users ==================================
func (ur *UserRepository) Get() ([]u.User, error) {
	users := []u.User{}
	if err := ur.db.Find(&users).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return users, nil
}

// ======================== Update User ==================================
func (ur *UserRepository) Update(userId int, newUser u.User) (u.User, error) {

	var user u.User
	ur.db.First(&user, userId)

	if err := ur.db.Model(&user).Updates(&newUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

// ======================== Delete User ==================================
func (ur *UserRepository) Delete(userId int) error {

	var user u.User

	if err := ur.db.First(&user, userId).Error; err != nil {
		return err
	}
	ur.db.Delete(&user, userId)
	return nil

}
// ============================================================================