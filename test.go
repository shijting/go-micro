package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shijting/go-micro/src/boot"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func waitForReadyConf() (err error) {
	// 使用context进行超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 6)
	defer cancel()
	for  {
		select {
		case <- ctx.Done():// 充nacos中读取配置超时
			err = fmt.Errorf("loaded config timeout")
			return
		default:
			if boot.JTConfig.Data.Mysql !=nil { // 读取到配置
				return
			}
		}
	}
}
// 中间件 如果还没从nacos中读取到配置
func checkForReadyConf() gin.HandlerFunc {
	return func(c *gin.Context) {
		if boot.JTConfig.Data.Mysql ==nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "server config is loading",
			})
		} else {
			c.Next()
		}
	}
}

func main() {
	boot.BootInit()
	r := gin.Default()
	// 中间件 如果还没从nacos中读取到mysql配置
	r.Use(checkForReadyConf())
	errNotify := make(chan error)
	go func() {
		// 用于 读取配置 的超时控制
		err := waitForReadyConf()

		if err != nil {
			fmt.Println("abc", err)
			errNotify <- err
		}
	}()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": boot.JTConfig.Data.Mysql,
		})
		return
	})
	go func() {
		err := r.Run(":8000")
		if err != nil {
			errNotify <- err
		}
	}()
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		errNotify <- fmt.Errorf("%s", <-sig)
	}()
	getErr := <- errNotify
	fmt.Println(123, getErr)
	// todo 关闭服务
}