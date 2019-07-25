package handler

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	example "micro/getSession/proto/example"
	"micro/ihome/utils"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Print */
	beego.Info("GetSession api/v1.0/session")

	/* Init */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/* get */
	sessionID := req.SessionID
	sessionKey := sessionID + "name"

	/* Redis */
	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("getSession Redis: ",err)
		return err
	}

	/* <=Redis */
	value := r.Get(sessionKey)
	str, _ := redis.String(value, nil)

	// return
	rsp.Data = str

	return nil
}
