package models

type Response struct {
	Header Header `json:"header"`
	Data   interface{} `json:"data"`
}

const (
	ServerSuccessCode = 1000
	ServerSuccessDesc = "success"

	FileTypeFile = "file"
	FileTypeImage = "image"
	FileTypeAudio = "audio"
	FileTypeVideo = "video"
)

type Header struct {
	Code        int `json:"code"`
	Description string `json:"description"`
}

