package main

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world!")
}

type ServiceRecord struct {
	ServiceName     string `json:"service_name"`
	ServiceIP       string `json:"service_ip"`
	ServicePort     int    `json:"service_port"`
	HealthCheckPath string `json:"health_check_path"`
	HttpScheme      string `json:"http_scheme"`
}

func main() {
	web.Router("/api/v1/services", &MainController{})
	web.Router("/api/v1/services", &MainController{})
	web.Run()
}
