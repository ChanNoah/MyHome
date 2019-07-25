package handler

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	example "micro/getPin/proto/example"
	"micro/ihome/models"
	"micro/ihome/utils"
	"strconv"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetPin(ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Print */
	beego.Info("GetPin api/v1.0/smscode/:mobile")

	/* Init */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/* MySQL*/
	// orm
	o := orm.NewOrm()
	// db object
	user := models.User{Mobile: req.Mobile}
	// sql
	err := o.Read(&user)
	if err == nil {
		beego.Error("GetPin MySQL sql: ", err)
		rsp.Errno = utils.RECODE_USERONERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/* Redis */
	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_mysql_dbname)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetPin Redis: ", err)
		return nil
	}

	/* <=Redis */
	// get
	value := r.Get(req.Uuid)
	// string
	valueString, _ := redis.String(value, nil)
	// if
	if req.Text != valueString {
		beego.Error("GetPin <=Redis if: Captcha error")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/* randon */
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := seed.Intn(8999) + 1000
	beego.Info("Pin: ", pin)

	/* Sim */
	//messageconfig := make(map[string]string)
	//messageconfig["appid"] = "29672"
	//messageconfig["appkey"] = "89d90165cbea8cae80137d7584179bdb"
	//messageconfig["signtype"] = "md5"
	//
	//messagexsend := submail.CreateMessageXSend()
	//submail.MessageXSendAddTo(messagexsend, req.Mobile)
	//submail.MessageXSendSetProject(messagexsend, "NQ1J94")
	//submail.MessageXSendAddVar(messagexsend, "code", strconv.Itoa(pin))
	//send := submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig)
	//fmt.Println("MessageXSend: ", send)
	//
	///* verify */
	//bo := strings.Contains(send, "success")
	//if bo == false {
	//	fmt.Println("GetPin verify: success error")
	//	rsp.Errno = utils.RECODE_SMSERR
	//	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//	return nil
	//}

	/* Redis<= */
	err = r.Put(req.Mobile, strconv.Itoa(pin), time.Second*300)
	if err != nil {
		beego.Error("GetPin Redis<= Put: ", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	return nil
}
