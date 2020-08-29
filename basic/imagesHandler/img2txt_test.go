package imagesHandler

import "testing"

func TestImg2txt(t *testing.T) {
	imgPath := "wallhaven-r2o9lm.png"
	//txts := []string{"@", "#", "*", "%", "+", ",", ".", " "}
	txts := []string{"*", "%", "#", "@", ",", "-", "+", " "}
	rowed := "\n"
	output := "保存的图片文本1"
	Img2txt(imgPath, 200, txts, rowed, output)
}
