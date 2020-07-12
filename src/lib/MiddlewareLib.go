package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-micro/framework/gin_"
	"github.com/shijting/go-micro/src/boot"
)

// 通用中间件

// 检查服务是否可用
func CheckMiddleware() gin_.Middleware {
	return func(next gin_.Endpoint) gin_.Endpoint {
		return func(ctx *gin.Context, request interface{}) (response interface{}, err error) {
			if !boot.ServerIsReady() {
				return nil, fmt.Errorf("server is loading")
			}
			return next(ctx, request)
		}
	}
}
