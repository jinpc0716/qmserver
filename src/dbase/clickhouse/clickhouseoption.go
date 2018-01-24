package clickoption

import (
	"github.com/astaxie/beego"
	"qmserver/src/entitydefine"
	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse"
)

var (
	// 定义常量
	CONNECT       [100]*sqlx.DB
	CLICK_HOST    string
	CLICK_DB      int
	CLICK_TIMEOUT int
	CLICK_OPEN	  int
	CLICK_MAXIDLE int
	CURRENCOUNT   int
)


func ClickHouseInit(){
	
	// 从配置文件获取ClickHouse的ip以及db
	CLICK_HOST  = beego.AppConfig.String("ClickHouse::host")
	CLICK_DB, _ = beego.AppConfig.Int("ClickHouse::dbase")
	CLICK_TIMEOUT, _ = beego.AppConfig.Int("ClickHouse::timeout")
	CLICK_OPEN = beego.AppConfig.DefaultInt("ClickHouse::DBOpen",0)
	CLICK_MAXIDLE = beego.AppConfig.DefaultInt("ClickHouse::maxidle",0)
	CURRENCOUNT = 0 

	if CLICK_OPEN == 0 {
		return
	}
	//首先建立基础连接数
	for index := 0; index < CLICK_MAXIDLE; index++ {
		ClickHouseConn(index)
	}
}

func ClickHouseConn( num int) {
	var err error
	CONNECT[num], err = sqlx.Open("clickhouse", CLICK_HOST)
	if err != nil {
		beego.Debug(err)
	}
}


func ClickHouseSelect(arr *[]interface{},ExchangID string,InstrumentID string,StartDay string,EndDay string,mtype int) bool {
	
	if CLICK_OPEN == 0 {
		return false
	}
	CURRENCOUNT++
	SQL:="SELECT DISTINCT TradingDay,OpenPrice,HighestPrice,LowestPrice,ClosePrice,LastPrice,Volume,Turnover,InstrumentID,UpdateTime,ActionDay,OpenInterest  from  " +
	ExchangID+"."+InstrumentID+" "+"where ActionDay >='"+StartDay+"' and ActionDay <='"+EndDay+"'  Order by ActionDay,Number desc"

	var dataarr []MarketStruct.OneDataStructDB
	if err := CONNECT[CURRENCOUNT%CLICK_MAXIDLE].Select(&dataarr, SQL); err != nil {
		beego.Debug(err)
		return false
	}
	if len(dataarr) > 0 {
		*arr = append(*arr, dataarr)
	}
	return true
}
