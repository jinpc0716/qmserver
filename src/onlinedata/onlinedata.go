package onlinedata

import (
	"strings"
    "github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"qmserver/src/dbase/redis"
	"qmserver/src/dbase/dboption"
	"qmserver/src/unity"
	"qmserver/src/entitydefine"
	"qmserver/src/R"
	"strconv"
	"time"
)

var (
	// 定义常量
	ONLINERedisClient   [100]redis.Conn
	ONLINE_HOST    string
	ONLINE_DB      int
	INDEXID		   string
	// 先声明map
	ExchangMap 	   map[string]string
)


func KafkaInit(){

	//获取配置信息
	ONLINE_HOST  = beego.AppConfig.String("redis::host")
	ONLINE_DB, _ = beego.AppConfig.Int("redis::dbase")
	ExchangMap = make(map[string]string)
	//启动消费者
	KafkaNum,_ := beego.AppConfig.Int("KafkaTopicNum")
	for index := 0; index < KafkaNum; index++ {
	 	//生成连接
		ONLINERedisClient[index] = dbredisoption.RedisConnect(ONLINE_HOST)
		go Consumer(index+1)
	 }
}

func Consumer(indexID int) {

	stripaddress:="onlineonemin"+strconv.Itoa(indexID)+"::ipaddress"
	address:=beego.AppConfig.String(stripaddress)
	strtopicid:="onlineonemin"+strconv.Itoa(indexID)+"::topicid"
	topicid:=beego.AppConfig.String(strtopicid)
	//获取上次的偏移量
	dbredisoption.RedisBindDB(0,ONLINERedisClient[indexID-1]);
	offersetdb:=dbredisoption.RedisGetValueByKey(topicid,ONLINERedisClient[indexID-1])
	if offersetdb == "" {
		offersetdb="-1"
	}
	
	master, err :=sarama.NewConsumer([]string{address}, nil)
	if err != nil {
		beego.Debug("NewConsumer Error!!!")
	}

	int64offerset,err := strconv.ParseInt(offersetdb, 10, 64) 
	consumer, err := master.ConsumePartition(topicid, 0, int64offerset)
	if err != nil {
		beego.Debug("ConsumePartition Error!!!")
	}

	for  {
		select {
		case message := <-consumer.Messages():
			{
				//如果是已经收过了的，不再接收
				if message.Offset == int64offerset  {
					continue
				}

				offerset:=strconv.FormatInt(message.Offset,10)  
				var mdat MarketStruct.OneDataStructInner

				//数据解析
				strvar := strings.Split(string(message.Value), ",")
				if len(strvar) < 12{
					continue
				}
	
				unity.StringToKlineStruct(strvar,&mdat)

				//判断数据类型
				dtype,_ :=strconv.Atoi(strvar[0]) 

				beego.Debug(topicid,message.Offset,string(message.Key),strvar[0],strvar[1],strvar[2],
					strvar[3],strvar[4],strvar[5],strvar[6],strvar[7],strvar[8],
					strvar[9],strvar[10],strvar[11],strvar[12])
				if dtype == R.ONEMINDATAKLINE {	

					//分钟线存储
					for {
						if dboption.DbInsert(1,mdat) {
							break
						}
						time.Sleep( 1 * time.Second )
					} 
					 
				} else if dtype == R.DAYDATAKLINE {
					//日线存入 db1
					
				}else{
					beego.Debug("Out data type!!!!!")
					continue
				}
				//更新offerset偏移量
				dbredisoption.RedisBindDB(0,ONLINERedisClient[indexID-1]);
				dbredisoption.RedisSetKeyAndValue(topicid,offerset,ONLINERedisClient[indexID-1])
			}
		}
	}
	beego.Debug("Done consuming topic hello")
	consumer.Close()
}



// consumer 消费者、、redis模块
// func Consumer(indexID int) {

// 	stripaddress:="onlineonemin"+strconv.Itoa(indexID)+"::ipaddress"
// 	address:=beego.AppConfig.String(stripaddress);
// 	strtopicid:="onlineonemin"+strconv.Itoa(indexID)+"::topicid"
// 	topicid:=beego.AppConfig.String(strtopicid);
// 	//获取上次的偏移量
// 	dbredisoption.RedisBindDB(0,ONLINERedisClient[indexID-1]);
// 	offersetdb:=dbredisoption.RedisGetValueByKey(topicid,ONLINERedisClient[indexID-1])
// 	if offersetdb == "" {
// 		offersetdb="-1"
// 	}
	
// 	master, err :=sarama.NewConsumer([]string{address}, nil)
// 	if err != nil {
// 		beego.Debug("NewConsumer Error!!!")
// 	}

// 	int64offerset,err := strconv.ParseInt(offersetdb, 10, 64) 
// 	consumer, err := master.ConsumePartition(topicid, 0, int64offerset)
// 	if err != nil {
// 		beego.Debug("ConsumePartition Error!!!")
// 	}

// 	for  {
// 		select {
// 		case message := <-consumer.Messages():
// 			{
// 				//如果是已经收过了的，不再接收
// 				if message.Offset == int64offerset  {
// 					continue
// 				}

// 				offerset:=strconv.FormatInt(message.Offset,10)  
// 				var mdat MarketStruct.OneDataStructInner

// 				//数据解析
// 				strvar := strings.Split(string(message.Value), ",")
// 				if len(strvar) < 12{
// 					continue
// 				}

// 				unity.StringToKlineStruct(strvar,&mdat)

// 				//判断数据类型
// 				dtype,_ :=strconv.Atoi(strvar[0]) 

// 				//beego.Debug(topicid,message.Offset,string(message.Key),strvar[0],strvar[1],strvar[2],
// 				//	strvar[3],strvar[4],strvar[5],strvar[6],strvar[7],strvar[8],
// 				//	strvar[9],strvar[10],strvar[11],strvar[12])

// 				if dtype == R.ONEMINDATAKLINE {	
// 					dbredisoption.RedisBindDB(1,ONLINERedisClient[indexID-1]);
// 					dbredisoption.RedisInsert(mdat,dtype,ONLINERedisClient[indexID-1])
// 				} else if dtype == R.DAYDATAKLINE {
// 					//日线存入 db1
// 					dbredisoption.RedisBindDB(2,ONLINERedisClient[indexID-1]);
// 					dbredisoption.RedisInsert(mdat,dtype,ONLINERedisClient[indexID-1])
// 				}else{
// 					beego.Debug("Out data type!!!!!")
// 					continue
// 				}
// 				//更新offerset偏移量
// 				dbredisoption.RedisBindDB(0,ONLINERedisClient[indexID-1]);
// 				dbredisoption.RedisSetKeyAndValue(topicid,offerset,ONLINERedisClient[indexID-1])
// 			}
// 		}
// 	}
// 	beego.Debug("Done consuming topic hello")
// 	consumer.Close()
// }

