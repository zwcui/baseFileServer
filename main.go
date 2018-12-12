package main

import (
	"io/ioutil"
	"os"
	"net/http"
	"fmt"
	"log"
	"baseFileServer/handler"
	"github.com/astaxie/beego"
)


const VERSION = "1.0"

//var viewsDirectory string
//var UploadDirectory string
var port string
var logging bool

func init(){
	//viewsDirectory 		= beego.AppConfig.String("viewsDirectory")	//页面展示地址
	//UploadDirectory   	= beego.AppConfig.String("uploadDirectory")		//上传地址
	port 				= ":"+beego.AppConfig.String("httpport")	//端口
	logging				= true			//是否开启日志
}

/*
	1. 文件服务器，http://localhost:8088 访问页面进行上传，一次不超过10个文件
	2. 相关路径端口等配置在上方，docker部署注意路径的挂载
	3. 上传参数中有uri，则按指定路径上传
	4. .apk文件上传时不修改文件名，其他以uuid作为文件名
	5. 如果安装了ffmpeg，则可以转文件格式
 */
func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println("Version " + VERSION)
		os.Exit(0)
	}

	if logging == false {
		log.SetOutput(ioutil.Discard)
	}


	mux := http.NewServeMux()

	mux.Handle("/", addDefaultHeaders(http.HandlerFunc(handler.IndexHandler)))
	mux.Handle("/uploadFile", addDefaultHeaders(http.HandlerFunc(handler.FileUploadHandler)))

	log.Printf("Listening on port %s .....\n", port)
	http.ListenAndServe(port, mux)

}

//添加默认header，用于跨域
func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}

