package dboption

import (
	"qmserver/src/dbase/clickhouse"
	"qmserver/src/entitydefine"
	"github.com/astaxie/beego"
)

var (
	dbtype int
)

//DBTYPE 1表示用clickhouse存储  2表示用mysql存储
func Dbinit(){
	dbtype,_ = beego.AppConfig.Int("DBTYPE")
	if dbtype == 1  {
		//用作查询
		clickoption.ClickHouseInit();
		//用作插入
		clickoption.ClickHouse2Init();
	}
}

func DbSelect(arr *[]interface{},ExchangID string,InstrumentID string,StartDay string,EndDay string,mtype int) bool{
	if dbtype == 1  {
		return clickoption.ClickHouseSelect(arr,ExchangID,InstrumentID,StartDay,EndDay,mtype);
	}
	return false
}

func DbInsert(num int,data MarketStruct.OneDataStructInner) bool {
	if dbtype == 1  {
		return clickoption.ClickHouseInsertKline(num,data);
	}
	return true
}