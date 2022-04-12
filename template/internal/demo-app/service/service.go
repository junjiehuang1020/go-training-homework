package service

import "junjiehuang1020.me/go-training-homework/template/internal/demo-app/biz"

func NewUserService(biz *biz.UserBiz) *UserService {

	return &UserService{
		userBiz: biz,
	}
}
