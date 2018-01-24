package unity


//客户机
import (
	"strconv"
	"qmserver/src/entitydefine"
//	"github.com/astaxie/beego"
	
	"strings"
)

//吧tick文件里的数据转换成struct
func TickStringToStruct(strvar []string,mdat *MarketStruct.OneDataStruct)  {
	mdat.TradingDay=strvar[0]
	mdat.HighestPrice,_ =strconv.ParseFloat(strvar[1], 64) 
	mdat.LowestPrice,_ =strconv.ParseFloat(strvar[2], 64) 
	mdat.OpenPrice,_ =strconv.ParseFloat(strvar[3], 64) 
	mdat.ClosePrice,_ =strconv.ParseFloat(strvar[4], 64) 
	mdat.LastPrice,_ =strconv.ParseFloat(strvar[5], 64)     
	mdat.Volume,_=strconv.Atoi(strvar[6])     
	mdat.Turnover,_ =strconv.ParseFloat(strvar[7], 64) 
	mdat.InstrumentID =strvar[8]
	mdat.UpdateTime   =strvar[9]
	if len(strvar) > 10  {
		mdat.ActionDay    =strvar[10]
	}
	if len(strvar) > 11  {
		mdat.OpenInterest,_ =strconv.ParseFloat(strvar[11], 64) 
	}
	
}

//exchanginfo 设置
func  ExchangInfoStruct(strvar []string,Exchaninfo *MarketStruct.ExchangeInfo){
	Exchaninfo.Exchangid=strvar[0]
	Exchaninfo.Productid=strvar[1]
	Exchaninfo.DayType,_=strconv.Atoi(strvar[2]) 
	Exchaninfo.Volume,_=strconv.Atoi(strvar[3]) 
	for index := 0; index < Exchaninfo.Volume ; index++ {
		var time MarketStruct.TimeTradingDay
		time.Starttime=strvar[index*2+4]
		time.Endtime=strvar[index*2+5]
		Exchaninfo.Time = append(Exchaninfo.Time,time)
	}
}

//exchangOrderType 设置
func  ExchangOrderTypeStruct(strvar []string,Exchaninfo *MarketStruct.ExchangeOrderType){
	Exchaninfo.Exchangid=strvar[0]
	Exchaninfo.OrderType=strvar[1]
	Exchaninfo.MarketPrice,_=strconv.Atoi(strvar[2])
}


func TickToMinStructChange(mdtick *MarketStruct.MarketData,mdmin *MarketStruct.OneDataStruct) {

	mdmin.TradingDay	=mdtick.TradingDay
	mdmin.OpenPrice		=mdtick.OpenPrice
	mdmin.HighestPrice	=mdtick.HighestPrice
	mdmin.LowestPrice	=mdtick.LowestPrice
	mdmin.ClosePrice	=mdtick.ClosePrice
	mdmin.LastPrice		=mdtick.LastPrice
	mdmin.Volume 		=mdtick.Volume
	mdmin.Turnover		=mdtick.Turnover
	mdmin.InstrumentID	=mdtick.InstrumentID
	mdmin.UpdateTime	=mdtick.UpdateTime
}

//吧tick文件里的数据转换成struct
func StringToKlineStruct(strvar []string,mdat *MarketStruct.OneDataStructInner)  {
	mdat.TradingDay=strvar[1]
	mdat.ExchangeID=strvar[2]
	mdat.HighestPrice=strvar[3]
	mdat.LowestPrice=strvar[4]
	mdat.OpenPrice=strvar[5]
	mdat.ClosePrice=strvar[6]
	mdat.LastPrice=strvar[7]    
	mdat.Volume=strvar[8]    
	mdat.Turnover=strvar[9]
	str := strings.Replace(strvar[10], "&", "", -1)
	str = strings.Replace(str, "-", "", -1)
	str = strings.Replace(str, " ", "", -1)
	mdat.InstrumentID =str
	mdat.UpdateTime   =strvar[11]
	mdat.ActionDay    =strvar[12]
	if len(strvar) > 13  {
		mdat.OpenInterest    =strvar[13]
	}
}

//吧tick文件里的数据转换成struct
func KlineStructToString(strlist *string,mdat *MarketStruct.OneDataStructInner)  {
	*strlist+=mdat.TradingDay+","
	*strlist+=mdat.OpenPrice+","
	*strlist+=mdat.HighestPrice+","
	*strlist+=mdat.LowestPrice+","
	*strlist+=mdat.ClosePrice+","
	*strlist+=mdat.LastPrice+","
	*strlist+=mdat.Volume+","
	*strlist+=mdat.Turnover+","
	*strlist+=mdat.InstrumentID+","
	*strlist+=mdat.ExchangeID+","
	*strlist+=mdat.UpdateTime+","
	*strlist+=mdat.ActionDay+","
	*strlist+=mdat.OpenInterest
}

//吧tick文件里的数据转换成struct
func StringToApiKlineStruct(strvar []string,mdat *MarketStruct.OneDataStruct)  {
	mdat.TradingDay=strvar[0]
	mdat.OpenPrice,_ =strconv.ParseFloat(strvar[1], 64) 
	mdat.HighestPrice,_ =strconv.ParseFloat(strvar[2], 64) 
	mdat.LowestPrice,_ =strconv.ParseFloat(strvar[3], 64) 
	mdat.ClosePrice,_ =strconv.ParseFloat(strvar[4], 64) 
	mdat.LastPrice,_ =strconv.ParseFloat(strvar[5], 64)     
	mdat.Volume,_=strconv.Atoi(strvar[6])     
	mdat.Turnover,_ =strconv.ParseFloat(strvar[7], 64) 
	mdat.InstrumentID =strvar[8]
	mdat.UpdateTime   =strvar[10]
	mdat.ActionDay    =strvar[11]
	if len(strvar) > 12 {
		mdat.OpenInterest,_=strconv.ParseFloat(strvar[12], 64)
	}
}