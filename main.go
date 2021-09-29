package main

import "github.com/beego/beego/v2/server/web"

func main() {
	ctrl := &MainController{}
	web.Router("/hello", ctrl)

	web.Run("127.0.0.1:8089")
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

// address: http://localhost:8080 GET
func (ctrl *MainController) Get() {

	name := ctrl.GetString("name")
	if name == "" {
		ctrl.Ctx.WriteString("Hello World")
		return
	}
	ctrl.Ctx.WriteString("Hello " + name)
}
