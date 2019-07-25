package handler

import (
	"context"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"image/color"
	example "micro/getCaptcha/proto/example"
	"micro/ihome/utils"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetCaptcha(ctx context.Context, req *example.Request, rsp *example.Response) error {
	/* Print */
	beego.Info("GetCaptcha api/v1.0/imagecode/:uuid")

	/* Inil */
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/* Captcha */
	cap := captcha.New()

	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	cap.SetSize(90, 41)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	img, str := cap.Create(4, captcha.NUM)
	beego.Info("captcha: ", str)

	/* Response */
	image := *img
	rgba := *(image.RGBA)

	for _, value := range rgba.Pix {
		rsp.Pinx = append(rsp.Pinx, uint32(value))
	}

	rsp.Stride = int64(rgba.Stride)

	rsp.Min = &example.Point{X:int64(rgba.Rect.Min.X), Y:int64(rgba.Rect.Min.Y)}
	rsp.Max = &example.Point{X:int64(rgba.Rect.Max.X), Y:int64(rgba.Rect.Max.Y)}

	/* Redis */
	r, err := utils.RedisServer(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err!=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/* Redis<= */
	err = r.Put(req.Uuid, str, time.Second*300)
	if err != nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		beego.Error("GetCaptcha Redis<=: ", err)
		return nil
	}

	/* return */
	return nil
}
