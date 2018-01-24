package clickoption

import (
	"database/sql"
	"github.com/astaxie/beego"
	"qmserver/src/entitydefine"
	"github.com/kshvakov/clickhouse"
	"qmserver/src/unity"
	"strconv"
)

var (
	// 定义常量
	CONNECT_T       [100]*sql.DB
	CLICK_HOST_T    string
	CLICK_DB_T      int
	CLICK_OPEN_T	int
	CLICK_MAXIDLE_T int
	CURRENCOUNT_T   int
)


func ClickHouse2Init(){
	
	// 从配置文件获取ClickHouse的ip以及db
	CLICK_HOST_T   = beego.AppConfig.String("ClickHouse::host")
	CLICK_DB_T , _ = beego.AppConfig.Int("ClickHouse::dbase")
	CLICK_OPEN_T  = beego.AppConfig.DefaultInt("ClickHouse::DBOpen",0)
	CLICK_MAXIDLE_T  = beego.AppConfig.DefaultInt("ClickHouse::numwritelink",0)
	CURRENCOUNT_T  = 0 

	if CLICK_OPEN_T  == 0 {
		return
	}
	//首先建立基础连接数
	for index := 0; index < CLICK_MAXIDLE_T; index++ {
		ClickHouse2Conn(index)
	}
}

func ClickHouse2Conn( num int) {
	var err error
	CONNECT_T[num], err = sql.Open("clickhouse", CLICK_HOST)
	if err != nil {
		beego.Debug(err)
	}
	if err := CONNECT_T[num].Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			beego.Debug(exception.Code, exception.Message, exception.StackTrace)
		} else {
			beego.Debug(err)
		}
		return
	}
}

func ClickHouseInsertKline(num int,data MarketStruct.OneDataStructInner) bool{

	if  data.TradingDay == "" {
		return true
	}
	if CLICK_OPEN_T == 0 {
		return false
	}

	CURRENCOUNT_T++
	//年
	stryear:=unity.Substr(data.TradingDay,0,4)
	//月
	strmonth:=unity.Substr(data.TradingDay,4,2)
	//日
	strday:=unity.Substr(data.TradingDay,6,2)

	tx, err   := CONNECT_T[CURRENCOUNT_T%CLICK_MAXIDLE_T].Begin()
	if err != nil {
		beego.Debug(err)
	}
	//拆分字符串
	SQL:="INSERT INTO " + data.ExchangeID+"."+data.InstrumentID+" (Number,IndexDate,TradingDay,OpenPrice,HighestPrice,LowestPrice,ClosePrice,LastPrice,Volume,Turnover,InstrumentID,UpdateTime,ActionDay,OpenInterest) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	
	stmt,err1 := tx.Prepare(SQL)
	if err1 != nil {

		if exception, ok := err1.(*clickhouse.Exception); ok {
			//数据库不存在,需要创建数据库
			if exception.Code == 60 || exception.Code == 81 || exception.Code == 82 {
				CreateDataTable(CONNECT_T[CURRENCOUNT_T%CLICK_MAXIDLE_T],data.ExchangeID,data.InstrumentID,0)
				stmt,err1 = tx.Prepare(SQL)
			}else {
				beego.Debug(err1,exception.Code,exception.Message)
				return false
			}
			
		}else{
			beego.Debug(err1)
			return false
		}
	}
	

	v1,_  := strconv.ParseFloat(data.OpenPrice, 64)
	v2,_  := strconv.ParseFloat(data.HighestPrice, 64)
	v3,_  := strconv.ParseFloat(data.LowestPrice, 64)
	v4,_  := strconv.ParseFloat(data.ClosePrice, 64)
	v5,_  := strconv.ParseFloat(data.LastPrice, 64)
	v6,_  := strconv.Atoi(data.Volume)
	v7,_  := strconv.ParseFloat(data.Turnover, 64)
	v8,_  := strconv.ParseFloat(data.OpenInterest,64)
	date :=stryear+"-"+strmonth+"-"+strday

	_, err = stmt.Exec(1,date,data.TradingDay,v1,v2,v3,v4,v5,v6,v7,data.InstrumentID,data.UpdateTime,data.ActionDay,v8)

	if err != nil {
		beego.Debug(err)
		return false
	}

	if err := tx.Commit(); err != nil {
		beego.Debug(err)
		return false
	}
	return true
}

//创建表
func CreateDataTable(conn *sql.DB,ExchangID string,InstrumentID string,datatype int) bool {
	SQL:="CREATE TABLE "+ExchangID+"."+InstrumentID
	SQL=SQL+" (Number Int32,IndexDate Date,TradingDay String,OpenPrice Float64,HighestPrice Float64,LowestPrice Float64,ClosePrice Float64,LastPrice Float64,Volume Int32,Turnover Float64,InstrumentID   String,UpdateTime String,ActionDay String,OpenInterest Float64)ENGINE =  MergeTree(IndexDate, (IndexDate,Number), 8192)"
	if _, err := conn.Exec(SQL); err != nil {
		 if exception, ok := err.(*clickhouse.Exception); ok {
			 //数据库不存在建库
			if exception.Code == 81 {
				if CreateDataBase(conn,ExchangID,datatype) {
					return CreateDataTable(conn,ExchangID,InstrumentID,datatype)
				}
			}else{
				beego.Debug(err)
				return false
			}
		 } else {
			beego.Debug(err)
			return false
		}
	}
	return true
}

//创建数据库
func CreateDataBase(conn *sql.DB,ExchangID string,datatype int) bool {
	SQL:="CREATE DATABASE "+ExchangID+";"
	//SQL=SQL+"(Number Int32,IndexDate Date,TradingDay String,OpenPrice Float64,HighestPrice Float64,LowestPrice Float64,ClosePrice Float64,LastPrice Float64,Volume Int32,Turnover Float64,InstrumentID   String,UpdateTime String,ActionDay String,OpenInterest Float64)ENGINE =  MergeTree(IndexDate, (IndexDate,Number), 8192)"
	if _, err := conn.Exec(SQL); err != nil {
		// if exception, ok := err.(*clickhouse.Exception); ok {
			
		// } else {
			beego.Debug(err)
		//}
		return false
	}
	return true
}