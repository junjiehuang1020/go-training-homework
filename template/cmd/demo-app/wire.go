//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"junjiehuang1020.me/go-training-homework/template/internal/demo-app/biz"
	"junjiehuang1020.me/go-training-homework/template/internal/demo-app/data"
	"junjiehuang1020.me/go-training-homework/template/internal/demo-app/service"
)

func InitializeService() *service.UserService {

	wire.Build(
		service.NewUserService,
		biz.NewUserBiz,
		wire.Bind(new(data.UserRepository), new(*data.UserRepoImpl)),
		data.NewUserRepoImpl,
		data.NewDB)

	return &service.UserService{}

}
