package boot

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
)

const (
	CacheBasePath = "runtime/configcache"
	DataId        = "go-micro-mysql"
	Group         = "go-micro"
)

type DataConfig struct {
	Mysql *MySqlConfig
	Redis *RedisConfig
}
type GlobalConfig struct {
	Config *struct{
		Address string
		Path string
		Port uint64
	}
	Service *struct{
		Namespace string
		Name string
	}
	Data *DataConfig
}
var JTConfig *GlobalConfig
var nacosClient  config_client.IConfigClient
//初始化主文件（本地)配置 和初始化nacos连接
func InitConfig()  {
	// 加载本地配置
	//configFile:="application.yaml"
	configFile:="/go/go-micro/application.yaml"
	err:=config.LoadFile(configFile)
	if err!=nil{
		log.Fatal(err)
	}
	JTConfig =&GlobalConfig{Data: new(DataConfig)}

	err=config.Get("sjt").Scan(JTConfig)
	fmt.Println("config",JTConfig.Config)
	if err!=nil{
		log.Fatal(err)
	}

	// 初始化nacos连接
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      JTConfig.Config.Address,
			ContextPath: JTConfig.Config.Path,
			Port:        JTConfig.Config.Port,
		},
	}
	nacosClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err!=nil{
		log.Fatal("nacos init err",err )
	}

	//开始加载数据相关的配置
	JTConfig.Data.Mysql = new(MySqlConfig)

	// 监听mysql
	listenMysql(JTConfig.Data.Mysql, true)
	// TODO 监听redis
	// TODO 监听es
}
func listenMysql(model ConfigInterface, reload bool){
	err := os.MkdirAll(CacheBasePath, 0666)
	if err != nil {
		log.Fatal("mkdir for cache dir error,", err)
	}
	err = nacosClient.ListenConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
		OnChange: func(namespace, group, dataId, data string) { // 第一次加载也会执行OnChange

			fmt.Println("OnChange", data)
			shouldReload:=reload
			// 第一次加载 不执行Reload
			if  !model.IsLoad(){
				shouldReload=false  //如果model 没有被加载过,则不需要做重载
			}
			//fmt.Println("OnChange", data)
			// 模拟网络卡顿
			//time.Sleep(time.Second * 3)
			cacheFile:=fmt.Sprintf("%s/%s-%s.yaml", CacheBasePath, group,dataId)
			file,err:=os.OpenFile(cacheFile,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
			if err!=nil{
				log.Println(err)
				return
			}
			defer  file.Close()
			_,err=file.WriteString(data)
			if err!=nil{
				log.Println(err)
				return
			}
			err=config.LoadFile(cacheFile)
			if err!=nil{
				log.Println(err)
				return
			}
			err=config.Scan(model)
			if err!=nil{
				log.Println(err)
				return
			}
			if shouldReload{ //重载关键代码
				err:=model.Reload()
				if err!=nil{
					logger.Error(err)
					return
				}
				logger.Info(dataId," 重载完成")
			}
		},
	})
	if err!=nil{
		log.Println("listen config error,dataid:", DataId,err)
	}
}