package controllers

import (
	//"strconv"
	"qmserver/src/entitydefine"
	"github.com/astaxie/beego"
	"qmserver/src/httpbusiness"
)

// Operations about Users
type HisController struct {
	beego.Controller
}
// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /instrumentid [get]
func (u *HisController) Get() {

	var ifdata MarketStruct.InterFaceStruct  
	//获取变量
	ifdata.Datatype		=u.GetString("datatype")
	ifdata.Instrument	=u.GetString("instrumentid")
	ifdata.Startday		=u.GetString("startday")
	ifdata.Exchangeid	=u.GetString("exchangeid")
	ifdata.Daynum		=u.GetString("daynum")
	ifdata.Secretkey	=u.GetString("secretkey")
	ifdata.Rtnnum		=u.GetString("rtnnum")

	if ifdata.Secretkey == "" {
		ErrorDataRequest(u,"secretkey");
		return
	}
	if ifdata.Datatype == "" {
		ErrorDataRequest(u,"datatype");
        	return
	}
	if ifdata.Exchangeid == "" {
		ErrorDataRequest(u,"exchangeid");
		return
	}
	beego.Debug(ifdata.Exchangeid,ifdata.Instrument,ifdata.Datatype,ifdata.Startday,ifdata.Secretkey,ifdata.Daynum,ifdata.Rtnnum)
	Busniess.ReqRequestString(&ifdata);
	u.Data["json"] =ifdata.Arr
	u.ServeJSON()
	//u.Ctx.Output.Download("D:/SVNGROUP/QMARKET/branch/qmkserver/target/qmkserver/bin/productinfo.csv")
}

func ErrorDataRequest(u *HisController,msg string){
	u.Ctx.WriteString(msg+"is empty!!!")
	u.ServeJSON()
}

func (u *HisController) Post(){
	
	var ifdata MarketStruct.InterFaceStruct  
	//获取变量
	ifdata.Datatype		=u.GetString("datatype")
	ifdata.Instrument	=u.GetString("instrumentid")
	ifdata.Startday		=u.GetString("startday")
	ifdata.Exchangeid	=u.GetString("exchangeid")
	ifdata.Daynum		=u.GetString("daynum")
	ifdata.Secretkey	=u.GetString("secretkey")
	ifdata.Rtnnum		=u.GetString("rtnnum")

	if ifdata.Secretkey == "" {
		ErrorDataRequest(u,"secretkey");
		return
	}
	if ifdata.Datatype == "" {
		ErrorDataRequest(u,"datatype");
        	return
	}
	if ifdata.Exchangeid == "" {
		ErrorDataRequest(u,"exchangeid");
		return
	}

	beego.Debug(ifdata.Exchangeid,ifdata.Instrument,ifdata.Datatype,ifdata.Startday,ifdata.Secretkey,ifdata.Daynum,ifdata.Rtnnum)

	
	Busniess.ReqRequestString(&ifdata);
	u.Data["json"] =ifdata.Arr
	u.ServeJSON()


}

