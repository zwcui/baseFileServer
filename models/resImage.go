package models


type ResImage struct {
	ResFile
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (f *ResImage) AddAttribute(){
	f.ResFile.AddAttribute()
	// Image, Width & Height
	width, height, err := GetImageSize(f.Name)
	if err == nil {
		f.Width = width
		f.Height = height
	}
}

// Use imaging library to Get Image Size
func GetImageSize(name string) (int, int, error) {
	//img, err := imaging.Open(name)
	//if err == nil {
	//	srcBounds := img.Bounds()
	//	return srcBounds.Max.X, srcBounds.Max.Y, nil
	//}else {
	//	return 0, 0, err
	//}

	return 0, 0, nil
}
