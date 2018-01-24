package unity

import (
	"time"
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
	)


//日期增加
//日期增加，没有
func Timeadd(dyear int,dmonth int,dday int,area int)  string {

	mmonth:=IntToMonth(dmonth)
	baseTime := time.Date(dyear,mmonth,dday, 0, 0, 0, 0, time.UTC)
    date:= baseTime.AddDate(0 , 0, area )
    //字符串拼接
    var month,day,year string
	year=fmt.Sprintf("%04d",date.Year())
	month=fmt.Sprintf("%02d",date.Month())
	day=fmt.Sprintf("%02d",date.Day())
	//传出
    var strdate string
    strdate=year+month+day
 	return strdate
}

//日期增加
//日期增加，没有
func TimeaddForYear(dyear int,dmonth int,dday int,area int)  string {
    

        mmonth:=IntToMonth(dmonth)
        baseTime := time.Date(dyear-area,mmonth,dday, 0, 0, 0, 0, time.UTC)
        date:= baseTime.AddDate(0 , 0, area )
        //字符串拼接
        var month,day,year string
        year=fmt.Sprintf("%04d",date.Year())
        month=fmt.Sprintf("%02d",date.Month())
        day=fmt.Sprintf("%02d",date.Day())
        //传出
        var strdate string
        strdate=year+month+day
         return strdate
}

    
//时间增加
//logtype = 0 是一分增加
//logtype = 1 是一小时增加
//addcount 是增加因子
func TimeAlog(strdate string ,strtime string,logtype int,addconut int) (str1 string,str2 string) {

    Ttime1:=StringToAllTime(strdate,strtime,0)

    //beego.Debug(Ttime1)
    //分钟增加
    var retndate,retntime string
    if logtype==0 {
        Time2:=Ttime1.Add(time.Minute * time.Duration(addconut)) 
        year:=fmt.Sprintf("%04d",Time2.Year())
        month:=fmt.Sprintf("%02d",Time2.Month())
        day:=fmt.Sprintf("%02d",Time2.Day())
        hour:=fmt.Sprintf("%02d",Time2.Hour())
        minute:=fmt.Sprintf("%02d",Time2.Minute())
        sec:=fmt.Sprintf("%02d",0)

        retndate=year+month+day
        retntime=hour+":"+minute+":"+sec
    //小时增加
    }else if logtype==1 {
        Time2:=Ttime1.Add(time.Hour * time.Duration(addconut)) 
        year:=fmt.Sprintf("%04d",Time2.Year())
        month:=fmt.Sprintf("%02d",Time2.Month())
        day:=fmt.Sprintf("%02d",Time2.Day())
        hour:=fmt.Sprintf("%02d",Time2.Hour())
        minute:=fmt.Sprintf("%02d",Time2.Minute())
        sec:=fmt.Sprintf("%02d",0)

        retndate=year+month+day
        retntime=hour+":"+minute+":"+sec
    }

    //hour,min,sec := Time2.Clock()
    //beego.Debug(Time2) 
    return retndate,retntime
}

//日期的比较
func StrTimecmp(timeone string,timetwo string) bool {

	beego.Debug(timeone,timetwo)
	Otime := StringToTime(timeone)
	Ttime := StringToTime(timetwo)

	if  Otime.Equal(Ttime)  {
	  //处理逻辑
	  return false
	}else{
	  return true
	}

}

//月份int型转Month型
func IntToMonth( month int) time.Month {

	var mmonth time.Month
	switch month {
    case 1:
        mmonth=1
        break //可以添加
    case 2:
        mmonth=2
        break //可以添加
    case 3:
        mmonth=3
        break //可以添加
    case 4:
        mmonth=4
        break //可以添加
    case 5:
        mmonth=5
        break //可以添加
    case 6:
        mmonth=6
        break //可以添加
    case 7:
        mmonth=7
        break //可以添加
    case 8:
        mmonth=8
        break //可以添加
    case 9:
        mmonth=9
        break //可以添加
    case 10:
        mmonth=10
        break //可以添加
    case 11:
        mmonth=11
        break //可以添加
    case 12:
        mmonth=12
        break //可以添加
    default:
        mmonth=1
        break //可以添加
 	}
 	return mmonth

}

//字符串转时间型
func StringToTime( strtime string) time.Time {

	//年
	stryear:=Substr(strtime,0,4)
	//月
	strmonth:=Substr(strtime,4,2)
	//日
	strday:=Substr(strtime,6,2)

	iyear,_:=strconv.Atoi(stryear)
	imonth,_:=strconv.Atoi(strmonth)
	iday,_:=strconv.Atoi(strday)


	mmonth:=IntToMonth(imonth)
	rtntime := time.Date(iyear,mmonth,iday, 0, 0, 0, 0, time.UTC)
	
	return rtntime
    
}

//tyep = 0 年月日时分秒全部转换
//type = 1 不含秒
//type = 2 不含分秒
//type = 3 不含时分秒
func StringToAllTime( strdate string,strtime string,datetype int) time.Time {
    //年
    stryear:=Substr(strdate,0,4)
    //月
    strmonth:=Substr(strdate,4,2)
    //日
    strday:=Substr(strdate,6,2)

    //时
    ihour,_:=strconv.Atoi(Substr(strtime,0,2))
    //分
    imin,_:=strconv.Atoi(Substr(strtime,3,2))
    //秒
    //isec,_:=strconv.Atoi(Substr(strtime,6,2))

    iyear,_:=strconv.Atoi(stryear)
    imonth,_:=strconv.Atoi(strmonth)
    iday,_:=strconv.Atoi(strday)

    mmonth:=IntToMonth(imonth)

    rtntime := time.Date(iyear,mmonth,iday, ihour, imin, 0, 0, time.UTC)
    
    return rtntime  
}


//时间的比较
func StrAllTimecmp(dateone string,timeone string,datetwo string,timetwo string) bool {

    //beego.Debug(timeone,timetwo)
    oldtime := StringToAllTime(dateone,timeone,0)
    newtime := StringToAllTime(datetwo,timetwo,0)

    if  oldtime.Equal(newtime)  {
      //处理逻辑
        return false
    }else if oldtime.Before(newtime) {
        return false
    }else{
        return true
    }

}