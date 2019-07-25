package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"image"
	"image/png"
	deleteSession "micro/deleteSession/proto/example"
	getArea "micro/getArea/proto/example"
	getCaptcha "micro/getCaptcha/proto/example"
	getPin "micro/getPin/proto/example"
	getSession "micro/getSession/proto/example"
	getUserInfo "micro/getUserInfo/proto/example"
	"micro/ihome/models"
	"micro/ihome/utils"
	postLogin "micro/postLogin/proto/example"
	postRegister "micro/postRegister/proto/example"
	"net/http"
	"regexp"
)

//templete
/*
func ExampleCall(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// print
	beego.Info("PostLogin /api/v1.0/sessions")

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// grpc
	client:=grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.template", client.Client())
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// Content-Type:application/json
	w.Header().Set("Content-Type","application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
*/

func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("GetArea api/v1.0/areas")

	// grpc
	client := grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := getArea.NewExampleService("go.micro.srv.getArea", client.Client())
	rsp, err := exampleClient.GetArea(context.TODO(), &getArea.Request{

	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// get data
	areas := []models.Area{}
	for _, value := range rsp.Data {
		temp := models.Area{Id: int(value.Id), Name: value.Name}
		areas = append(areas, temp)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   areas,
	}

	// Content-Type:application/json
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// print
	beego.Info("GetIndex /api/v1.0/house/index")

	//

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),
	}

	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// print
	beego.Info("GetSession api/v1.0/session")

	// cookie
	cookie, err := r.Cookie("cookie")
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_DBERR),
		}

		// Content-Type:application/json
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// grpc
	client := grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := getSession.NewExampleService("go.micro.srv.getSession", client.Client())
	rsp, err := exampleClient.GetSession(context.TODO(), &getSession.Request{
		SessionID: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]string)
	data["name"] = rsp.Data

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	// Content-Type:application/json
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetCaptcha(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/* Print */
	beego.Info("GetCaptcha api/v1.0/imagecode/:uuid")

	/* Grpc */
	client := grpc.NewService()
	/* Inil */
	client.Init()

	/* Call */
	// 1 call the backend service
	exampleClient := getCaptcha.NewExampleService("go.micro.srv.getCaptcha", client.Client())
	rsp, err := exampleClient.GetCaptcha(context.TODO(), &getCaptcha.Request{
		Uuid: params.ByName("uuid"),
	})
	// 2 if
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if rsp.Errno != "0" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  rsp.Errno,
			"errmsg": rsp.Errmsg,
		}

		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	/* Response */
	// 1 rgba
	var rgba image.RGBA

	// pinx
	for _, value := range rsp.Pinx {
		rgba.Pix = append(rgba.Pix, uint8(value))

	}
	// stride
	rgba.Stride = int(rsp.Stride)

	// rectangle
	rgba.Rect.Min.X = int(rsp.Min.X)
	rgba.Rect.Min.Y = int(rsp.Min.Y)
	rgba.Rect.Max.X = int(rsp.Max.X)
	rgba.Rect.Max.Y = int(rsp.Max.Y)

	// 2 image
	var image captcha.Image

	// rgba
	image.RGBA = &rgba

	// encode
	png.Encode(w, image)
}

func GetPin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// print
	beego.Info("GetPin api/v1.0/smscode/:mobile")

	/* Get */
	// mobile
	mobile := params.ByName("mobile")

	MyRegexp := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)

	bo := MyRegexp.MatchString(mobile)

	if bo == false {
		response := map[string]interface{}{
			"errno":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}

		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	// url
	text := r.URL.Query()["text"][0] //captcha
	id := r.URL.Query()["id"][0]

	// grpc
	client := grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := getPin.NewExampleService("go.micro.srv.getPin", client.Client())
	rsp, err := exampleClient.GetPin(context.TODO(), &getPin.Request{
		Mobile: mobile,
		Text:   text,
		Uuid:   id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// srv -> web
	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func PostRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// print
	beego.Info("PostRegister api/v1.0/users")

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// if
	if request["mobile"] == "" || request["password"] == "" || request["pin"] == "" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}

		// Content-Type:application/json
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// grpc
	client := grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := postRegister.NewExampleService("go.micro.srv.postRegister", client.Client())
	rsp, err := exampleClient.PostRegister(context.TODO(), &postRegister.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		Pin:      request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	// cookie
	_, err = r.Cookie("cookie")

	if err != nil {
		cookie := http.Cookie{Name: "cookie", Value: rsp.SessionID, Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
	}

	// Content-Type:application/json
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// print
	beego.Info("PostLogin /api/v1.0/sessions")

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// get
	if request["mobile"] == "" || request["password"] == "" {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}

		// Content-Type:application/json
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// grpc
	client := grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := postLogin.NewExampleService("go.micro.srv.postLogin", client.Client())
	rsp, err := exampleClient.PostLogin(context.TODO(), &postLogin.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = r.Cookie("cookie")
	if err != nil {
		cookie := http.Cookie{Name: "cookie", Value: rsp.SessionID, Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	// Content-Type:application/json
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// print
	beego.Info("DeleteSession /api/v1.0/session")

	// cookie
	cookie, err := r.Cookie("cookie")
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		// Content-Type:application/json
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// grpc
	client := grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := deleteSession.NewExampleService("go.micro.srv.deleteSession", client.Client())
	rsp, err := exampleClient.DeleteSession(context.TODO(), &deleteSession.Request{
		SessionID: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// cookie
	if cookie.Value != "" {
		cookie := http.Cookie{Name: "cookie", Path: "/", MaxAge: -1}
		http.SetCookie(w, &cookie)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	// Content-Type:application/json
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// print
	beego.Info("GetUserInfo /api/v1.0/user")

	// cookie
	cookie, err := r.Cookie("cookie")
	if err != nil {
		// we want to augment the response
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		// Content-Type:application/json
		w.Header().Set("Content-Type", "application/json")

		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// grpc
	client := grpc.NewService()
	// inil
	client.Init()

	// call the backend service
	exampleClient := getUserInfo.NewExampleService("go.micro.srv.getUserInfo", client.Client())
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &getUserInfo.Request{
		SessionID: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// data to front-end
	data := make(map[string]string)

	data["user_id"] = rsp.UserId
	data["name"] = rsp.Name
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = rsp.AvatarUrl

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data": data,
	}

	// Content-Type:application/json
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
