package handler

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"micro/ihome/models"
	"micro/ihome/utils"
	"strconv"
	"time"

	example "micro/postLogin/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostLogin (ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Print */
	beego.Info("PostLogin /api/v1.0/sessions")

	/* Init */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/* MySQL*/
	// 1 orm
	o := orm.NewOrm()
	// 2 db object
	user := models.User{}
	// 3 select
	user.Mobile = req.Mobile
	err := o.Read(&user,"Mobile")
	if err != nil {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostLogin MySQL: ", err)
		return nil
	}

	// Md5
	hash, err := utils.CrotoMd5(req.Password)
	if err != nil{
		rsp.Errno = utils.RECODE_UNKNOWERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostLogin Md5: ", err)
		return nil
	}
	// password
	if hash != user.Password_hash{
		rsp.Errno = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostLogin password: ", err)
		return nil
	}

	/* sessionID */
	rsp.SessionID, err = utils.CrotoMd5(req.Mobile + req.Password + strconv.Itoa(int(time.Now().UnixNano())))
	if err != nil {
		rsp.Errno = utils.RECODE_UNKNOWERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostLogin sessionID: ", err)
		return nil
	}

	/* Redis */
	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("PostLogin Redis")
		return nil
	}

	/* <=Redis */
	err = r.Put(rsp.SessionID+"name", user.Name, time.Second*600)
	err = r.Put(rsp.SessionID+"id", user.Id, time.Second*600)
	err = r.Put(rsp.SessionID+"mobile", user.Mobile, time.Second*600)

	return nil
}

