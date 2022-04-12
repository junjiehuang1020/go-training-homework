package data

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	QueryUserByid(uid int32) (UserPO, error)
}

type UserRepoImpl struct {
	db gorm.DB
}

func (u *UserRepoImpl) QueryUserByid(uid int32) (UserPO, error) {

	var user UserPO

	u.db.Where("id = ?", uid).First(&user)

	return user, nil

}

type UserPO struct {
	Id       int32
	Name     string
	Gender   int32
	Age      int32
	Address  string
	Level    int32
	CreateAt time.Time
	UpdateAt time.Time
}
