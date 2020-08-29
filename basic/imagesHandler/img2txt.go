package imagesHandler

import (
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func Img2txt(imgPath string, size uint, txts []string, rowed, output string)  {
	// 获取图片文件
	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println("打开图片文件失败: ", err.Error())
		return
	}
	defer file.Close()

	// 用图片文件获取图片对象
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("获取图片对象失败：", err.Error())
		return
	}

	// 将宽度设置为size,然后换算出等比例的高度
	var width = size
	var height = (size * (uint(img.Bounds().Dy()))) / (uint(img.Bounds().Dx()))
	height = height * 6 / 10 // 6/10是大致字符的宽高比
	newing := resize.Resize(width, height, img, resize.Lanczos3)	// //根据高宽resize图片，并得到新图片的像素值
	dx := newing.Bounds().Dx()
	dy := newing.Bounds().Dy()

	// 创建一个字节buffer,用来存放字符
	textBuffer := bytes.Buffer{}
	// 字符切片长度
	txtLen := len(txts)
	// 遍历图片每一行每一列像素
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			colorRGB := newing.At(x, y)
			r, g, b, _ := colorRGB.RGBA()

			// 获取三原色值，取平均数
			avg := uint8((r+g+b)/3 >> 8)

			//有多少个用来替换的字符就将256分为多少个等分，
			//然后计算这个像素的平均值趋紧与哪个字符，最后，将这个字符添加到字符buffer里
			num := avg / uint8(256/txtLen)
			textBuffer.WriteString(txts[num])
			fmt.Print(txts[num])
		}

		textBuffer.WriteString(rowed)	// 一行结束，换行
		fmt.Print(rowed)
	}

	// 将字符buffer的数据存入文本，结束
	f, err := os.Create(output + ".txt")
	if err != nil {
		fmt.Println("创建文本文件失败:", err.Error())
		return
	}
	defer f.Close()

	f.WriteString(textBuffer.String())
}
