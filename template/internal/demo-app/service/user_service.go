package service

import (
	"junjiehuang1020.me/go-training-homework/template/internal/demo-app/biz"
)

type UserDTO struct {
	Name    string `json:"name"`
	Gender  int32  `json:"gender"`
	Age     int32  `json:"age"`
	Address string `json:"address"`
	Level   int32  `json:"level"`
}

type UserService struct {
	userBiz *biz.UserBiz
}

func (s *UserService) UserInfo(uid int32) (*UserDTO, error) {

	do, error := s.userBiz.GetUserById(uid)
	if error != nil {
		//todo
	}
	return &UserDTO{
		Name:    do.Name,
		Gender:  do.Gender,
		Age:     do.Age,
		Address: do.Address,
		Level:   do.Level,
	}, nil

}
