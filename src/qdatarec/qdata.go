package qdp

import (
	"qmserver/src/dbase/redis"
	"qmserver/struct/qdp"
)
//数据量合约查询
func InstrumentSelect(instrument string) (qdpstruct.MarketData,int) {
		temdata,err := dbase.Dbread(instrument)
		if err != -1 {
				return temdata,0
		}	
		return temdata,-1
}

//实时行情查询，先去map中寻找，如果找不到再去数据库寻找，最后找不到返回err -1
func LastPriceSearch(instrument string,exchangeid string ) (qdpstruct.MarketData,int){
	var data qdpstruct.MarketData
	if exchangeid !="" {
		if data, ok := Datalist[exchangeid][instrument]; ok {
			return data,0
		}
	}else{
		for _,value:=range Datalist{
			data,ok := value[instrument]
			if ok {
				return data,0
			}
		}
		//map里没有去数据库中寻找，找到以后复制给map
		data,err := InstrumentSelect(instrument)
		if err != -1 {
			mm, ok := Datalist[exchangeid]
			if  !ok  {
				mm = make(map[string]qdpstruct.MarketData)
	    		Datalist[exchangeid] = mm
			}
			Datalist[exchangeid][instrument]= data
			return data,0
		}
	}
	return data,-1
}