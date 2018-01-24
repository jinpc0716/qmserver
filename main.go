package main

import (
	"qmserver/src/httpbusiness"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"qmserver/controllers"
	"qmserver/src/onlinedata"
	//"qmserver/src/qdprec"
	"qmserver/src/dbase/redis"
	//"qmserver/src/unity"
	"qmserver/src/dbase/dboption"
)

func main() {

	//历史行情入口
	beego.Router("/hismin", &controllers.HisController{})

	//查询参数初始化
	Busniess.InitReqRequestInfo()

	//历史数据库初始化
	dboption.Dbinit()

	//初始化缓存数据库
	dbredisoption.Redisinit()
	
	//kafka数据接受程序
	go onlinedata.KafkaInit()

	//日志输出目录
	logpath:=beego.AppConfig.String("qmserver::logpath");
	var console_config=`{"filename":"`+logpath+`","level":7}`
	logs.SetLogger(logs.AdapterFile, console_config)

	//日志异步输出
	logs.Async()
	beego.Debug("main start")

	beego.Run()
}