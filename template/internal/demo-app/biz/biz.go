package biz

import (
	"junjiehuang1020.me/go-training-homework/template/internal/demo-app/data"
)

func NewUserBiz(userRepo data.UserRepository) *UserBiz {
	return &UserBiz{
		userRepo: userRepo,
	}
}
