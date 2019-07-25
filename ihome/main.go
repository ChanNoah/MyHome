package main

import (
        "github.com/julienschmidt/httprouter"
        "github.com/micro/go-log"
        "github.com/micro/go-web"
        "micro/ihome/handler"
        "net/http"

        _ "micro/ihome/models"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.IhomeWeb"),
                web.Version("latest"),
                // address
                web.Address("127.0.0.1:10086"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

    // new router
    router:=httprouter.New()

    // register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))
    router.NotFound = http.FileServer(http.Dir("html"))

	// register call handler
	//service.HandleFunc("/", handler.ExampleCall)
    //router.GET("/example/call",handler.ExampleCall)
    router.GET("/api/v1.0/areas",handler.GetArea)
    router.GET("/api/v1.0/house/index",handler.GetIndex)
    router.GET("/api/v1.0/imagecode/:uuid",handler.GetCaptcha)
    router.GET("/api/v1.0/smscode/:mobile",handler.GetPin)
    router.POST("/api/v1.0/users",handler.PostRegister)
    router.GET("/api/v1.0/session",handler.GetSession)
    router.POST("/api/v1.0/sessions",handler.PostLogin)
    router.DELETE("/api/v1.0/session",handler.DeleteSession)
    router.GET("/api/v1.0/user",handler.GetUserInfo)


    service.Handle("/", router)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
