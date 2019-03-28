package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/savitarai/mon-server-go/models"
)

type RecordController struct {
	beego.Controller
}

func (r *RecordController) Post() {
	var report models.RecordRequest
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &report)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", report)
	r.Data["json"] = map[string]string{"status": "ok"}
	r.ServeJSON()
}
