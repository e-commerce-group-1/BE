package user

import (
	"errors"
	u "group-project1/entities/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// ======================== User Register ==================================
func (ur *UserRepository) Insert(newUser u.Users) (u.Users, error) {
	if err := ur.db.Create(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

// ======================== Get Users ==================================
func (ur *UserRepository) Get() ([]u.Users, error) {
	users := []u.Users{}
	ur.db.Find(&users)
	if len(users) < 1 {
		return nil, errors.New("belum ada user yang terdaftar")
	}
	return users, nil
}

// ======================== Update User ==================================
func (ur *UserRepository) Update(userUpdate u.Users) (u.Users, error) {
	res := ur.db.Model(&userUpdate).Updates(userUpdate)
	if res.RowsAffected == 0 {
		return userUpdate, errors.New("tidak ada pemutakhiran pada data user")
	}
	ur.db.First(&userUpdate)
	return userUpdate, nil
}

// ======================== Delete User ==================================
func (ur *UserRepository) Delete(userId int) error {
	var user u.Users
	res := ur.db.Delete(&user, userId)
	if res.RowsAffected == 0 {
		return errors.New("tidak ada user yang dihapus")
	}
	return nil

}

// ============================================================================
