package MarketStruct


//交易所信息
type ExchangeInfo struct{

	Productid	string
	Exchangid	string
	DayType 	int		//交易节类型1 表示是当日日期，2表示是次日日期
	Volume 		int
	Time     [] TimeTradingDay
}

//开盘收盘时间 交易节
type TimeTradingDay struct{
	Starttime string
	Endtime   string
}

type InterFaceStruct struct{
	Datatype		string
	Instrument      string
	Startday		string
	Exchangeid		string
	Daynum			string
	Secretkey		string
	Rtnnum			string
	Arr []interface{}
}

//交易所订单类型
type ExchangeOrderType struct{
	Exchangid	string
	OrderType 	string	//GFD;FAK;FOK
	MarketPrice int  //1表示支持市价单
}
