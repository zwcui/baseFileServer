package handler

import (
	"net/http"
	"html/template"
	"github.com/astaxie/beego"
)

type home struct {
	Title string
}

//首页handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	title := home{Title: "首页"}
	t, _ := template.ParseFiles(beego.AppConfig.String("viewsDirectory")+  "upload.html")
	t.Execute(w, title)
}

