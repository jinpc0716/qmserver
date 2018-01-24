package tests

import (
	"qmserver/src/dbase"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
)

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

func testredis(){

	rc:=dbredisoption.RedisClient.Get()
	defer rc.Close()
	username, err := redis.String(rc.Do("GET", "ilanni"))
    if err != nil {
        beego.Debug("redis get failed:", err)
    } else {
		beego.Debug("redis get failed:", username)
    }
}

// func TestShort(t *testing.T) {
// 	request := beetest.Post("/v1/shorten")
// 	request.Param("longurl", "http://www.beego.me/")
// 	response, _ := request.Response()
// 	defer response.Body.Close()
// 	contents, _ := ioutil.ReadAll(response.Body)
// 	var s ShortResult
// 	json.Unmarshal(contents, &s)
// 	if s.UrlShort == "" {
// 		t.Fatal("shorturl is empty")
// 	}
// }

// func TestExpand(t *testing.T) {
// 	request := beetest.Get("/v1/expand")
// 	request.Param("shorturl", "5laZF")
// 	response, _ := request.Response()
// 	defer response.Body.Close()
// 	contents, _ := ioutil.ReadAll(response.Body)
// 	var s ShortResult
// 	json.Unmarshal(contents, &s)
// 	if s.UrlLong == "" {
// 		t.Fatal("urllong is empty")
// 	}
// }
