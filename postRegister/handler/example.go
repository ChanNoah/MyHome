package handler

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"micro/ihome/models"
	"micro/ihome/utils"
	"time"

	example "micro/postRegister/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostRegister(ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Print */
	beego.Info("PostRegister api/v1.0/users")

	/* Init */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/* Redis */
	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostRegister Redis: ", err)
		return nil
	}

	/* <=Redis */
	// get
	value := r.Get(req.Mobile)
	// string
	valueString, _ := redis.String(value, nil)
	// if
	if valueString != req.Pin{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostRegister <=Redis if: Pin error")
	}

	/* Croto */
	str, err := utils.CrotoMd5(req.Password)
	if err != nil{
		rsp.Errno = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostRegister Croto: ", err)
		return nil
	}

	/* MySQL */
	// orm
	o := orm.NewOrm()
	// db object
	user := models.User{}
	// SQL
	user.Password_hash = str
	user.Name = req.Mobile
	user.Mobile = req.Mobile
	// insert
	id, err := o.Insert(&user)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostRegister MySQL insert: ", err)
		return nil
	}

	/* session */
	rsp.SessionID, err = utils.CrotoMd5(req.Mobile+string(time.Now().UnixNano()))

	/* Redis<= */
	err = r.Put(rsp.SessionID+"name", user.Name, time.Second*300)
	err = r.Put(rsp.SessionID+"mobile", user.Mobile, time.Second*300)
	err = r.Put(rsp.SessionID+"id", id, time.Second*300)

	/* Return */
	return nil
}
