package main

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/yanyiwu/gojieba"
)

func main() {
	ctrl := &MainController{}
	jiebaCtrl := &JiebaController{}

	web.Router("/hello", ctrl)
	web.Router("/jieba", jiebaCtrl)

	web.Run("127.0.0.1:8089")
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

type JiebaController struct {
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

var x = gojieba.NewJieba()

func (jiebaCtrl *JiebaController) Get() {

	sentence := jiebaCtrl.GetString("sentence")

	var words []string
	words = x.Cut(sentence, true)

	jiebaCtrl.Ctx.WriteString(strings.Join(words, "/"))
}
