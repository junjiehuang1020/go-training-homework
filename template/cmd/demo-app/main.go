package main

import "github.com/gin-gonic/gin"

func handler(ctx *gin.Context) {

	params := struct {
		Msg string `json:"msg"`
	}{}

	ctx.BindQuery(&params)

	ctx.JSON(200, gin.H{
		"message": params.Msg,
	})

}

func main() {

	r := gin.Default()
	r.GET("/ping", handler)
	r.Run()

}
