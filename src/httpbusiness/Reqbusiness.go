package Busniess

import (
	//"shorturl/models"
	"qmserver/src/entitydefine"
	"qmserver/src/dataOption"
	"strconv"
	"qmserver/src/R"
	"github.com/astaxie/beego"
	"strings"
	//"qmserver/src/dbase/clickhouse"
)

var (
	// 定义常量
	ExchangInfoPath    		string
	ExchangOrderTypePath    string
	DataFilePath 			string
)

//初始化读历史数据信息
func InitReqRequestInfo(){
	//交易节信息目录
	ExchangInfoPath=beego.AppConfig.String("qmserver::ExchangInfoPath")
	//获取交所交易指令类型目录
	ExchangOrderTypePath=beego.AppConfig.String("qmserver::ExchangOrderTypePath")
	//数据文件
	DataFilePath=beego.AppConfig.String("qmserver::datapath");
}
	

//业务层分析 传入参数属于是什么类型 要做什么业务查询 。另外密钥验证也在这里做
//先做密钥验证再处理
func ReqRequestString(ifdata *MarketStruct.InterFaceStruct){
	
	//登录认证

	//判断是读取什么
	//datatype=0 一分值  ，1 日线值  2- 8 保留 999交易所开盘收盘时间查询
	dtype,_:=strconv.Atoi(ifdata.Datatype)
	switch dtype {
	case R.EXCHANGINFO:
		{
			//获取交所交易节信息
			CsvOption.ReadExchangInfoToArr(ExchangInfoPath,ifdata.Exchangeid,&ifdata.Arr)
			break
		}
	case R.EXCHANGORDERTYPE:
		{
			//获取交所交易指令类型
			CsvOption.ReadExchangOrderTypeToArr(ExchangOrderTypePath,ifdata.Exchangeid,&ifdata.Arr)
			break
		}
	case R.ONEMINDATA,R.DAYLINEDATA:
		{
			if ifdata.Startday == "" {
				beego.Debug("startday is null!!");
				return
			}
			if ifdata.Instrument == "" {
				beego.Debug("Instrument is null!!");
				return
			}
			//如果是黄金 且 是(T D)
			if (strings.Contains(ifdata.Instrument," ") && ifdata.Exchangeid == "SGE" ) {
				//替换空格
				ifdata.Instrument=strings.Replace(ifdata.Instrument," ","+",1);
			}
			dnum,_:=strconv.Atoi(ifdata.Daynum)
			rtnnum,_:=strconv.Atoi(ifdata.Rtnnum)
			CsvOption.OneMinDataReadString(&ifdata.Arr,ifdata.Exchangeid,ifdata.Startday,ifdata.Instrument,ifdata.Datatype,dnum,rtnnum,DataFilePath)
			break
		}
	case R.TICKDATA:
		{

			break
		}
	default:{
			beego.Debug("ERROR datatype!!!!")
		}
	}
}
