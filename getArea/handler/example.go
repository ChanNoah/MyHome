package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"micro/ihome/models"
	"micro/ihome/utils"
	"time"

	example "micro/getArea/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Pritn */
	beego.Info("GetArea api/v1.0/areas")

	/* Inil */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	// 2 db object
	areas := []models.Area{}

	/* Redis */
	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err!=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetArea Redis: ",err)
		return nil
	}

 	/* <=Redis */
 	// 1 get
 	area_json := r.Get("areas")

 	// 2 if
 	if area_json != nil{
 		// 2.1 unmarshal
 		json.Unmarshal((area_json).([]byte), &areas)

 		// 2.2 response
		for _, value := range areas {
			temp := example.Address{Id:int32(value.Id),Name:value.Name}
			rsp.Data = append(rsp.Data, &temp)
		}
		// 2.3 return
		return nil
	}

	/* Mysql */
	// 1 orm
	o:=orm.NewOrm()
	// 2 db object
	//areas := []models.Area{}
	// 3 sql
	qs:=o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetArea Mysql 3 sql", err)
		return  nil
	}
	if num == 0{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetArea Mysql 3 sql num = 0")
		return  nil
	}
	// 4 response
	for _, value := range areas {
		temp := example.Address{Id:int32(value.Id),Name:value.Name}
		rsp.Data = append(rsp.Data, &temp)
	}

	/* Redis<= */
	// 1 marshal
	areas_json,_ := json.Marshal(areas)

	// 2 put
	r.Put("areas",areas_json, time.Second * 3600)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetArea Redis<= 2 put", err)
		return nil
	}

	/* return */
	return nil
}
