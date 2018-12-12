package handler

import (
	"net/http"
	"fmt"
	"io"
	"strings"
	"path/filepath"
	"os"
	"log"
	"encoding/json"
	"github.com/satori/go.uuid"
	"baseFileServer/models"
	"github.com/astaxie/beego"
)

var dir string

func init(){
	dir = beego.AppConfig.String("uploadDirectory")		//上传地址

	//路径如果是.，就取当前工程路径
	if dir == "." {
		dir, _ = filepath.Abs(dir)
		log.Println(dir)
		dir += "/data/"
		log.Println(dir)
	}

	if _, err := os.Stat(dir); err != nil {
		log.Printf("Directory %s not exist, Create it", dir)
		errPath := os.MkdirAll(dir, 0777)
		if errPath != nil {
			log.Fatalf("Directory %s not exist, Create it Fail", dir)
			return
		}
	}

	//补全路径
	if strings.HasSuffix(dir, "/") == false {
		dir = dir + "/"
	}
}

//文件上传handler
//上传参数中有uri，则按指定路径上传
//.apk文件上传时不修改文件名，其他以uuid作为文件名
//如果安装了ffmpeg，则可以转文件格式
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fileArray []interface{}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		randomUUId, _ := uuid.NewV4()
		paSplit := strings.Split(randomUUId.String(), "-")
		// File Path
		var relatedFileDir string
		if r.URL.Query().Get("uri") != "" {
			relatedFileDir = r.URL.Query().Get("uri")
		}else {
			relatedFileDir = paSplit[0] + "/" + paSplit[1] + "/" + paSplit[2] + "/" + paSplit[3] + "/"
		}

		fileExt := filepath.Ext(part.FileName())
		fileName := ""
		if strings.EqualFold(".apk", strings.ToLower(fileExt)) {
			fileName = part.FileName()
		} else {
			// File Name
			fileName = paSplit[4] + filepath.Ext(part.FileName())
		}


		// Create File Dir if not
		var fileDir string
		if dir != "." {
			fileDir = dir + relatedFileDir
		}
		errPath := os.MkdirAll(fileDir, 0777)
		if errPath != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create File
		ff := fileDir + fileName
		dst, err := os.Create(ff)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Print("receive file successfully! fileName:" + part.FileName())
		
		//if strings.EqualFold(r.URL.Query().Get("fileType"), "audio") {
		//	models.ChangeVoice(r.URL.Query().Get("changeFlag"), relatedFileDir, fileName)
		//	var fileSuffix string
		//	fileSuffix = path.Ext(fileName)
		//	var filenameOnly string
		//	filenameOnly = strings.TrimSuffix(fileName, fileSuffix)
		//	fileName = filenameOnly + ".mp3"
		//}
		// append resFile to response Array
		fileArray = append(fileArray, models.ResFileFromFileName(dst.Name(), "/" + relatedFileDir + fileName, r.URL.Query().Get("fileType")))

		if len(fileArray) >= 10 {
			// 最多上传10个文件
			break;
		}
	}

	res := models.Response{Header:models.Header{Code:models.ServerSuccessCode, Description:models.ServerSuccessDesc}, Data:fileArray}

	// Generate Json
	jsonByte, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write Response
	fmt.Fprint(w, string(jsonByte))
	return
}


