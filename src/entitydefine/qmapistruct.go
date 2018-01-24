package MarketStruct


type ApiData struct {
	TradingDay         string 
	PreSettlementPrice string
	PreClosePrice      string
	PreOpenInterest    string
	OpenPrice          string
	HighestPrice       string
	LowestPrice        string
	ClosePrice         string
	UpperLimitPrice    string
	LowerLimitPrice    string
	SettlementPrice    string
	CurrDelta          string
	LastPrice          string
	Volume             string
	Turnover           string
	OpenInterest       string
	BidPrice1          string
	BidVolume1         string  
	AskPrice1          string
	AskVolume1         string  
	BidPrice2          string
	BidVolume2         string  
	BidPrice3          string
	BidVolume3         string  
	AskPrice2          string
	AskVolume2         string  
	AskPrice3          string
	AskVolume3         string  
	BidPrice4          string
	BidVolume4         string 
	BidPrice5          string
	BidVolume5         string  
	AskPrice4          string
	AskVolume4         string  
	AskPrice5          string
	AskVolume5         string  
	InstrumentID       string  `orm:"pk"`
	UpdateTime         string 
	UpdateMillisec     string
	NowVolume          string 
	ExchangeId     	   string
	InstrumentName 	   string
}

