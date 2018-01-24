package CsvOption

//客户机
import (
	"strconv"
	"qmserver/src/dbase/redis"
	"qmserver/src/entitydefine"
	"github.com/astaxie/beego"
	"qmserver/src/R"
	"qmserver/src/unity"
	"encoding/csv"
	"io"
	"os"
)

//k线数据读取循环遍历取文件
func OneMinDataRead( arr *[]MarketStruct.OneDataStruct,exchangeid string,startday string,instrument string,datatype string ) {
	
	//年
	stryear:=unity.Substr(startday,0,4)

	dtype,err:=strconv.Atoi(datatype);
	if err != nil {
		return
	}

	HisDataPath:=beego.AppConfig.String("qmserver::datapath");

	var strpath string
	if dtype == R.ONEMINDATA {
		//月
		strmonth:=unity.Substr(startday,4,2)
		//日
		strpath=HisDataPath+exchangeid+"_onemin"+"/"+stryear+"/"+strmonth+"/"+instrument+"_"+startday+".csv"	
	}else{
		//日
		strpath=HisDataPath+exchangeid+"_day"+"/"+stryear+"/"+instrument+"_"+stryear+".csv"
	}

	//var strdate string
	if unity.Exist(strpath) {
		ReadCsvToStruct(strpath,arr)
		//strdate=unity.Timeadd(iyear,imonth,iday,1)
		//loopDate(arr,exchangeid,strdate,endday,instrument)
	}else{	
		beego.Debug("文件不存在！！%s",strpath);
	}
	
}


//k现数据读取字符串转换结构体
func ReadCsvToStruct( path string,arr *[]MarketStruct.OneDataStruct )  {
	
	//给了文件路径和map的结构体
	file, err := os.Open(path)
    if err != nil {
        return
    }
	defer file.Close()

	var flagfirstline int32

    reader := csv.NewReader(file)
    for {

		flagfirstline++
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            return
		}

		if flagfirstline == 1  {
			continue
		}
		
		var mdat *MarketStruct.OneDataStruct = new(MarketStruct.OneDataStruct)
		unity.TickStringToStruct(record,mdat) 
		*arr = append(*arr, *mdat)
		//data, _ := json.Marshal(mdat)
    }
}

func OneMinDataReadFromDB( arr *[]interface{},exchangeid string,startday string,instrument string,datatype string,
	allday int,rtnnum int)  {
	//遍历交易日历
	
	//判断要取哪几天数据，然后通过数据库来获取这几天数据
	//startday endday
	
}

//k线数据读取循环遍历取文件
func OneMinDataReadString( arr *[]interface{},exchangeid string,startday string,instrument string,datatype string,
	allday int,rtnnum int,HisDataPath string )  int {

	//年
	stryear:=unity.Substr(startday,0,4)
	//月
	strmonth:=unity.Substr(startday,4,2)
	//日
	strday:=unity.Substr(startday,6,2)

	dtype,err:=strconv.Atoi(datatype);
	if err != nil {
		return -1
	}

	iyear,_:=strconv.Atoi(stryear)
	imonth,_:=strconv.Atoi(strmonth)
	iday,_:=strconv.Atoi(strday)

	var strpath string
	if dtype == R.ONEMINDATA {

		if rtnnum > 20  {
			beego.Warning("rtnnum request rtn num too long!!")
			return -1
		}
		if rtnnum < 0  {
			beego.Warning("rtnnum request rtn num too short!!")
			return -1
		}
	
		if allday > 4  {
			beego.Warning("allday request day too long!!")
			return 0
		}
	
		if allday <= 0  {
			return 0
		}

		//月
		strmonth:=unity.Substr(startday,4,2)
		//日
		strpath=HisDataPath+"/onemin/"+exchangeid+"/"+stryear+"/"+strmonth+"/"+instrument+"_"+startday+".csv"
		
		beego.Debug(strpath)
		tbool:=dbredisoption.RedisSelect(arr,dtype,exchangeid,startday,instrument)
		if(!tbool){
			//daynum 0 表示当日， n表示往后取n天的数据最大不能超过5天
			if unity.Exist(strpath) {
				ReadCsvToString(strpath,arr)
				strdate:=unity.Timeadd(iyear,imonth,iday,-1)
				OneMinDataReadString(arr,exchangeid,strdate,instrument,datatype,allday-1,rtnnum,HisDataPath)
				//往后再获取n天数据
				return 0
			}else{
				strdate:=unity.Timeadd(iyear,imonth,iday,-1)
				OneMinDataReadString(arr,exchangeid,strdate,instrument,datatype,allday,rtnnum-1,HisDataPath)
				return 0
			}
		}else{
			strdate:=unity.Timeadd(iyear,imonth,iday,-1)
			OneMinDataReadString(arr,exchangeid,strdate,instrument,datatype,allday,rtnnum-1,HisDataPath)
			return 0
		}

	}else{

		if rtnnum > 5  {
			beego.Warning("rtnnum request rtn num too long!!")
			return -1
		}

		if rtnnum < 0  {
			beego.Warning("rtnnum request rtn num too short!!")
			return -1
		}

		if allday > 3  {
			beego.Warning("allday request day too long!!")
			return 0
		}

		if allday <= 0  {
			return 0
		}

		//日
		strpath=HisDataPath+"/day/"+exchangeid+"/"+stryear+"/"+instrument+"_"+stryear+".csv"
		beego.Debug(strpath)
		if unity.Exist(strpath) {
			ReadCsvToString(strpath,arr)
			strdate:=unity.TimeaddForYear(iyear,imonth,iday,1)
			OneMinDataReadString(arr,exchangeid,strdate,instrument,datatype,allday-1,rtnnum,HisDataPath)
			return 0
		}else{
			strdate:=unity.TimeaddForYear(iyear,imonth,iday,1)
			OneMinDataReadString(arr,exchangeid,strdate,instrument,datatype,allday,rtnnum-1,HisDataPath)
			return 0
		}
	}
}


//k现数据读取字符串转换结构体
func ReadCsvToString(path string, arr *[]interface{})  {
	
	//给了文件路径和map的结构体
	file, err := os.Open(path)
    if err != nil {
        return
    }
	defer file.Close()

	var flagfirstline int32
	flagfirstline=0
	var arrtmp []MarketStruct.OneDataStruct
	reader := csv.NewReader(file)
    for {
		flagfirstline++
        record, err := reader.Read()
        if err == io.EOF {
            break
        }

		if flagfirstline == 1  {
			continue
		}
		var mdat MarketStruct.OneDataStruct 
		unity.TickStringToStruct(record,&mdat) 
		arrtmp = append(arrtmp, mdat)
	}
	index := 0
	for  ; index < len(arrtmp) ; index++ {
		*arr = append(*arr, arrtmp[len(arrtmp) - index-1] )
	}
}

//读取交易所信息
func ReadExchangInfoToArr(path string,exchange_id string,arr *[]interface{})   {
	
	//给了文件路径和map的结构体
	file, err := os.Open(path)
    if err != nil {
        return
	}
	
	defer file.Close()
	var flagfirstline int

	reader := csv.NewReader(file)
    for {

		flagfirstline++
        record, err := reader.Read()
        if err == io.EOF {
			break
		}
        // } else if err != nil {
        //     return
		// }
		if flagfirstline == 1  {
			continue
		}
		var ExchangeInfo MarketStruct.ExchangeInfo 
		//匹配交易所 如果是相同的交易所就传输
		if exchange_id == "*"	{

			unity.ExchangInfoStruct(record,&ExchangeInfo)
			*arr = append(*arr, &ExchangeInfo)

		}else if record[0] == exchange_id  {

			unity.ExchangInfoStruct(record,&ExchangeInfo)
			*arr = append(*arr, &ExchangeInfo)
		}
	}
}



//读取交易所信息
func ReadExchangOrderTypeToArr(path string,exchange_id string,arr *[]interface{})   {
	
	//给了文件路径和map的结构体
	file, err := os.Open(path)
    if err != nil {
        return
	}
	
	defer file.Close()
	var flagfirstline int

	reader := csv.NewReader(file)
    for {

		flagfirstline++
        record, err := reader.Read()
        if err == io.EOF {
			break
		}
        // } else if err != nil {
        //     return
		// }
		if flagfirstline == 1  {
			continue
		}
		var ExchangeInfo MarketStruct.ExchangeOrderType 
		//匹配交易所 如果是相同的交易所就传输
		if exchange_id == "*"	{
			unity.ExchangOrderTypeStruct(record,&ExchangeInfo)
			*arr = append(*arr, &ExchangeInfo)
		}else if record[0] == exchange_id  {
			unity.ExchangOrderTypeStruct(record,&ExchangeInfo)
			*arr = append(*arr, &ExchangeInfo)
		}
	}
}