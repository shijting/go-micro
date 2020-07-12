package boot

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
)


//mysql相关
var mysqlDB *gorm.DB

func WaitForDbReady(d time.Duration)  {
	go func() {
		err:=WaitForReady(d,func() error {
			return InitMysql()
		},"数据库初始化成功","数据库初始化失败")
		if err!=nil{
			ErrChan<-err
		}
	}()

}
func ReloadDB() error  {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	if mysqlDB!=nil{
		err:=mysqlDB.Close()
		if err!=nil{
			return err
		}
		mysqlDB=nil
		return InitMysql()
	}
	return nil

}
func InitMysql() error {
	//fmt.Println("hhhh", JTConfig.Data.Mysql.Dsn)
	var err error
	mysqlDB, err = gorm.Open("mysql",
		JTConfig.Data.Mysql.Dsn)
	if err != nil {
		mysqlDB = nil
		 return NewFatalError(err.Error()) //这里返回致命异常
	}
	mysqlDB.SingularTable(true)
	mysqlDB.DB().SetMaxIdleConns(JTConfig.Data.Mysql.Maxidle)
	mysqlDB.DB().SetMaxOpenConns(JTConfig.Data.Mysql.Maxopen)
	return nil
}
func GetDB() *gorm.DB  {
	return mysqlDB
}
