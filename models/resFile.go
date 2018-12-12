package models
import (
	"os"
	"strings"
	"os/exec"
	"path"
	"log"
)

type ResFileWrapper interface {
	// Add File Attribute
	AddAttribute()
}

type ResFile struct {
	Name     string `json:"-"`
	Uri      string `json:"uri"`
	Size     int64 `json:"size"`
	FileType string `json:"fileType"`
}

func (f *ResFile) AddAttribute() {
	file, err := os.Open(f.Name)
	if err != nil {
		return
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		f.Size = stat.Size()
	}
}


// Construct ResFile
func ResFileFromFileName(name string, uri string, fileType string) interface{} {
	var res ResFileWrapper
	resFile := ResFile{Name:name, Uri:uri, FileType:fileType}

	switch fileType {
	//case FileTypeImage:
	//	res = &ResImage{ResFile:resFile}
	//case FileTypeAudio:
	//	res = &ResAudio{ResFile:resFile}
	//case FileTypeVideo:
	//	res = &ResVideo{ResAudio:ResAudio{ResFile:resFile}}
	default:
		res = &resFile
	}
	res.AddAttribute()
	return res
}

func ChangeVoice(changeFlag string, relatedFileDir string, fileName string){
	var fileSuffix string
	fileSuffix = path.Ext(fileName)
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(fileName, fileSuffix)
	var err error
	if strings.EqualFold(changeFlag, "change") {
		if strings.EqualFold(fileSuffix, ".wav") {
			//_, err := exec.Command("sh", "-c", "sox " + "data/" + relatedFileDir + fileName + " data/" + relatedFileDir + "change.wav pitch 300").Output()
			//if err == nil {
			//	_, _ = exec.Command("sh", "-c", "lame -h  data/" + relatedFileDir + "change.wav" + " data/" + relatedFileDir + filenameOnly + ".mp3").Output()
			//	_, _ = exec.Command("sh", "-c", "rm -rf data/" + relatedFileDir + "change.wav").Output()
			//	_, _ = exec.Command("sh", "-c", "rm -rf data/" + relatedFileDir + fileName).Output()
			//}
			_, err = exec.Command("sh", "-c", "ffmpeg -i " + "data/" + relatedFileDir + fileName + " -f mp3 data/" + relatedFileDir + filenameOnly + ".mp3").Output()
			//if err == nil {
			//	_, _ = exec.Command("sh", "-c", "rm -rf data/" + relatedFileDir + fileName).Output()
			//}
		} else if strings.EqualFold(fileSuffix, ".mp3") {
			//_, err := exec.Command("sh", "-c", "sox " + "data/" + relatedFileDir + fileName + " data/" + relatedFileDir + "change.wav pitch 300").Output()
			//if err == nil {
			//	_, _ = exec.Command("sh", "-c", "lame -h  data/" + relatedFileDir + "change.wav" + " data/" + relatedFileDir + fileName).Output()
			//	_, _ = exec.Command("sh", "-c", "rm -rf data/" + relatedFileDir + "change.wav").Output()
			//}
		}
	} else {
		switch fileSuffix {
		case ".wav" :
			_, err = exec.Command("sh", "-c", "ffmpeg -i " + "data/" + relatedFileDir + fileName + " -f mp3 data/" + relatedFileDir + filenameOnly + ".mp3").Output()
			//if err == nil {
			//	_, err = exec.Command("sh", "-c", "rm -rf data/" + relatedFileDir + fileName).Output()
			//}
		case ".amr" :
			_, err = exec.Command("sh", "-c", "ffmpeg -i " + "data/" + relatedFileDir + fileName + " -f mp3 data/" + relatedFileDir + filenameOnly + ".mp3").Output()
			//if err == nil {
			//	_, err = exec.Command("sh", "-c", "rm -rf data/" + relatedFileDir + fileName).Output()
			//}
		}
	}
	if err != nil {
		log.Print("音频文件err:" + err.Error())
	}
}
