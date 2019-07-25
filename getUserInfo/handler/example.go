package handler

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"micro/ihome/models"
	"micro/ihome/utils"
	"strconv"

	example "micro/getUserInfo/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Print */
	beego.Info("GetUserInfo api/v1.0/user")

	/* Init */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/* Get */
	sessionID := req.SessionID

	/* Redis */
	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetUserInfo Redis: ",err)
		return nil
	}

	/* <=Redis */
	key := sessionID+"id"
	value := r.Get(key)
	valueStr, err := redis.String(value,nil)
	if err!=nil{
		rsp.Errno = utils.RECODE_UNKNOWERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetUserInfo <=Redis redis.String: ",err)
		return nil
	}

	valueInt, err := strconv.Atoi(valueStr)
	if err!=nil{
		rsp.Errno = utils.RECODE_UNKNOWERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetUserInfo <=Redis strconv.Atoi: ",err)
		return nil
	}

	/* MySQL */
	// orm
	o := orm.NewOrm()

	// db object
	user:=models.User{}

	// select
	user.Id = valueInt
	err = o.Read(&user, "Id")
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetUserInfo MySQL select: ",err)
		return nil
	}

	// return
	rsp.UserId = string(user.Id)
	rsp.Name = user.Name
	rsp.Mobile = user.Mobile
	rsp.RealName = user.Real_name
	rsp.IdCard = user.Id_card
	rsp.AvatarUrl = user.Avatar_url

	/* Return */
	return nil
}

