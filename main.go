package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"newbooking/pkg/config"
)

func main() {
	r := gin.Default()

	config.RouterConfig(r)

	err := r.Run(":8888")
	if err != nil {
		fmt.Println(err.Error())
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
