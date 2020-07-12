package boot

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/logger"
	"time"
)

var ErrChan chan error
//启动函数，用于各种初始化，首先要完成配置中心的初始化
func BootInit()  {
	ErrChan =make(chan error)
	InitConfig()
	go func() {
		// 基础配置初始化
		err:=WaitForConfigReady(time.Second*5)
		if err!=nil{
			ErrChan <-err
		}
		//数据库初始化
		WaitForDbReady(time.Second*5)
	}()

}
//服务器已经准备好了 ,后面要不断扩展
func ServerIsReady() bool  {
	if  configIsReady() && mysqlDB!=nil{
		return true
	}
	return false
}
//通用的 超时控制 函数
func WaitForReady(d time.Duration,f func() error,success string,fail string) error {
	ctx,cancel:=context.WithTimeout(context.Background(),d)
	defer cancel()
	for{
		select {
		case <-ctx.Done():
			return fmt.Errorf(fail)
		default:
			err:=f()
			if err==nil{  //没有错误 则直接返回, 退出循环
				logger.Info(success)
				return nil
			}else if IsFatalError(err){  //如果是致命异常，退出循环
				return  fmt.Errorf("出现致命异常:%s",err.Error())
			}
		}
	}
}