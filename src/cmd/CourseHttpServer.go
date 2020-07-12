package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"github.com/shijting/go-micro/framework/gin_"
	"github.com/shijting/go-micro/src/boot"
	_ "github.com/shijting/go-micro/src/lib"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)
// 课程 http api
func main()  {
	// 加载配置
	boot.BootInit()
	r:=gin.New()
	gin_.BootStrap(r)
	// todo 408 timeout 错误
	service := web.NewService(
		web.Name(strings.Join([]string{boot.JTConfig.Service.Namespace, boot.JTConfig.Service.Name}, ".")),
		web.Handler(r),
		web.Address(":8888"),
	)

	go func() {
		if err := service.Init(); err !=nil {
			boot.ErrChan <- err
		}
		if err :=service.Run(); err !=nil {
			boot.ErrChan <- err
		}
	}()
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		boot.ErrChan <- fmt.Errorf("%s", <-sig)
	}()
	getErr := <- boot.ErrChan
	fmt.Println(getErr)
	log.Fatal(getErr)

}
