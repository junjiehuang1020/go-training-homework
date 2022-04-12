package biz

import (
	"errors"

	"junjiehuang1020.me/go-training-homework/template/internal/demo-app/data"
)

type UserBiz struct {
	userRepo data.UserRepository
}

func (biz *UserBiz) GetUserById(uid int32) (*UserDO, error) {

	if uid == 0 {
		return nil, errors.New("invalid user id")
	}

	user, err := biz.userRepo.QueryUserByid(uid)

	if err != nil {
		// 处理下错误一般可以返回一个错误码
	}

	return &UserDO{
		Id:      user.Id,
		Name:    user.Name,
		Gender:  user.Gender,
		Age:     user.Age,
		Address: user.Address,
		Level:   user.Level,
	}, nil

}

type UserDO struct {
	Id      int32
	Name    string
	Gender  int32
	Age     int32
	Address string
	Level   int32
}
