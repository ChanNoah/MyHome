package handler

import (
	"context"
	"github.com/astaxie/beego"
	"micro/ihome/utils"

	example "micro/deleteSession/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) DeleteSession (ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Print */
	beego.Info("DeleteSession /api/v1.0/session")

	/* Init */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/* Redis */

	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err!=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/* Redis<=*/
	err = r.Delete(req.SessionID+"name")
	err = r.Delete(req.SessionID+"id")
	err = r.Delete(req.SessionID+"mobile")

	return nil
}
