package dbredisoption

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"qmserver/src/entitydefine"
	"qmserver/src/unity"
	"qmserver/src/R"
	"time"
	"strings"
)

var (
	// 定义常量
	RedisClient   *redis.Pool
	REDIS_HOST    string
	REDIS_DB      int
	REDIS_TIMEOUT int
	REDIS_OPEN	  int
	REDIS_STATUE  bool
)

//建立连接池
func Redisinit() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST  = beego.AppConfig.String("redis::host")
	REDIS_DB, _ = beego.AppConfig.Int("redis::dbase")
	REDIS_TIMEOUT, _ = beego.AppConfig.Int("redis::timeout")
	REDIS_OPEN = beego.AppConfig.DefaultInt("redis::RedisOpen",0)
	REDIS_STATUE =true
	if REDIS_OPEN == 0 {
		return
	}
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		//MaxIdle：最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle:     beego.AppConfig.DefaultInt("redis::maxidle", 40),
		//MaxActive：最大的激活连接数，表示同时最多有N个连接
		MaxActive:   beego.AppConfig.DefaultInt("redis::maxactive", 100),
		//IdleTimeout：最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 200 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				REDIS_STATUE=false
				beego.Debug(err)
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

//建立库固定操作
func RedisConnect(host string)   redis.Conn {
	c, err := redis.Dial("tcp", host)
	if err != nil {
		REDIS_STATUE=false
		beego.Debug(err)
		return nil
	}
	return c
}
//选择数据库
func RedisBindDB(dbnum int,c redis.Conn)   redis.Conn {
	if !REDIS_STATUE {
		return nil
	}
	c.Do("SELECT", dbnum)
	return c
}

func RedisInsert(onedata MarketStruct.OneDataStructInner,datatype int,c redis.Conn) {
	if !REDIS_STATUE {
		return 
	}
	var strbuf string
	var instrument string
	unity.KlineStructToString(&strbuf,&onedata)
	if datatype == R.DAYDATAKLINE {
		//按照交易所+合约
		instrument=onedata.ExchangeID+onedata.InstrumentID
	}else{
		//按照交易所+合约+交易日设定key
		instrument=onedata.ExchangeID+onedata.InstrumentID+onedata.ActionDay
	}

	_, err := c.Do("GET",instrument )
    if err != nil {
        //保存列表
		_, err = c.Do("lpush",instrument,strbuf )
		if err != nil {
			beego.Error("RedisInsert error",err)
		}

		//设置过期时间，一个月的过期时间
		_, err = c.Do("expire",instrument, REDIS_TIMEOUT *24*60*60 )
		if err != nil {
			beego.Error("RedisInsert error",err)
		}

	}else{

		//保存列表
		_, err = c.Do("lpush",instrument,strbuf)
		if err != nil {
			beego.Error("RedisInsert error",err)
		}
	}

}

func RedisSetKeyAndValue(key string,value string,c redis.Conn){
	if !REDIS_STATUE {
		return 
	}
	_, err := c.Do("SET",key,value )
	if err != nil {
		beego.Error("RedisSetKeyAndValue error",err)
	}
}

func RedisGetValueByKey(key string,c redis.Conn) string {
	if !REDIS_STATUE {
		return ""
	}
	username, err := redis.String(c.Do("GET", key))
    if err != nil {
        return ""
	}
	return username
}


func RedisSelect(arr *[]interface{},dtype int,exchangeid string,startday string,instrument string) bool {

	if REDIS_OPEN == 0 {
		return false
	}

	var rtnbool bool
	rtnbool=false

	var instrumenttmp string
	//按照交易所+合约+交易日设定key
	if dtype == 0 {
		instrumenttmp =exchangeid+instrument
	}else{
		instrumenttmp =exchangeid+instrument+startday
	}
	
	rc:=RedisClient.Get()
	defer rc.Close()
	if rc == nil {
		beego.Error("Error ,rc is nill")
		return false
	}
	values,error := redis.Values(rc.Do("lrange",instrumenttmp,0,-1))
	if error != nil {
		beego.Error("RedisSelect error",error)
		rtnbool=false
	}else{
		for _, v := range values {
			var mdat MarketStruct.OneDataStruct
			//beego.Debug("RedisSelect",string(v.([]byte)))
			strvar := strings.Split(string(v.([]byte)), ",")
			unity.StringToApiKlineStruct(strvar,&mdat) 
			*arr = append(*arr, mdat)
			//数据转换
			rtnbool=true
		}
	}
	return rtnbool
}


