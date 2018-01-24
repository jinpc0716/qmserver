package MarketStruct

type MarketData struct {
	TradingDay         string 
	SettlementGroupID  string 
	SettlementID       int  
	PreSettlementPrice float64
	PreClosePrice      float64
	PreOpenInterest    float64
	PreDelta           float64
	OpenPrice          float64
	HighestPrice       float64
	LowestPrice        float64
	ClosePrice         float64
	UpperLimitPrice    float64
	LowerLimitPrice    float64
	SettlementPrice    float64
	CurrDelta          float64
	LastPrice          float64
	Volume             int
	Turnover           float64
	OpenInterest       float64
	BidPrice1          float64
	BidVolume1         int  
	AskPrice1          float64
	AskVolume1         int  
	BidPrice2          float64
	BidVolume2         int  
	BidPrice3          float64
	BidVolume3         int  
	AskPrice2          float64
	AskVolume2         int  
	AskPrice3          float64
	AskVolume3         int  
	BidPrice4          float64
	BidVolume4         int 
	BidPrice5          float64
	BidVolume5         int  
	AskPrice4          float64
	AskVolume4         int  
	AskPrice5          float64
	AskVolume5         int  
	InstrumentID       string  `orm:"pk"`
	UpdateTime         string 
	UpdateMillisec     int
	NowVolume          int  
	ExchangeId     	   string

}

type OneDataStruct struct {
	TradingDay         string
	OpenPrice          float64
	HighestPrice       float64
	LowestPrice        float64
	ClosePrice         float64
	LastPrice          float64
	Volume             int
	Turnover           float64 
	InstrumentID       string //`json:"onstrumentID"`
	UpdateTime         string  
	ActionDay          string 
	OpenInterest       float64 
}


type OneDataStructInner struct {
	TradingDay         string  `db:"TradingDay"`
	OpenPrice          string  `db:"OpenPrice"`
	HighestPrice       string  `db:"HighestPrice"`
	LowestPrice        string  `db:"LowestPrice"`
	ClosePrice         string  `db:"ClosePrice"`
	LastPrice          string  `db:"LastPrice"`
	Volume             string  `db:"Volume"`
	Turnover           string  `db:"Turnover"`
	InstrumentID       string  `db:"InstrumentID"`
	ExchangeID         string  `db:"ExchangeID"`
	UpdateTime         string  `db:"UpdateTime"`
	ActionDay          string  `db:"ActionDay"`
	OpenInterest       string  `db:"OpenInterest"`
}



type OneDataStructDB struct {
	TradingDay         string  `db:"TradingDay"`
	OpenPrice          float64  `db:"OpenPrice"`
	HighestPrice       float64  `db:"HighestPrice"`
	LowestPrice        float64  `db:"LowestPrice"`
	ClosePrice         float64  `db:"ClosePrice"`
	LastPrice          float64  `db:"LastPrice"`
	Volume             int32  `db:"Volume"`
	Turnover           float64  `db:"Turnover"`
	InstrumentID       string  `db:"InstrumentID"`
	UpdateTime         string  `db:"UpdateTime"`
	ActionDay          string  `db:"ActionDay"`
	OpenInterest       float64  `db:"OpenInterest"`
}

type Demo1 struct {
	UpdateTime     	   string 
	InstrumentName 	   string `orm:"pk"`
}


