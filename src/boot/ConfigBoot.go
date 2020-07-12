package boot

import (
	"fmt"
	"time"
)

func configIsReady() bool {

	//以下设定：哪些配置值必须有值，才算成功。有一个失败都算 not ready
	checks:=[]ConfigInterface{JTConfig.Data.Mysql}
	for _,c:=range checks{
		if !c.IsLoad(){
			return false
		}
	}
	//if JTConfig.Data.Mysql == nil {
	//	return false
	//}
	return true
}
//配置中心相关 初始化 ,,支持超时参数
func WaitForConfigReady(d time.Duration) error  {
	 return WaitForReady(d, func() error {
	 	res := configIsReady()
		 if res{
		 	return nil
		 }else{
		 	return fmt.Errorf("wait")
		 }
	 },"配置加载成功","配置加载失败")
}