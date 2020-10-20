package sync

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Mosaic(w http.ResponseWriter, r *http.Request)  {
	t0 := time.Now()
	// 从 POST 表单获取用户上传内容
	r.ParseMultipartForm(10485760)
	// 获取上传图片
	file, _, _ := r.FormFile("image")
	defer file.Close()
	// 读取用户设置的区块尺寸
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	// 解码原始图片
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	// 克隆目标图片用于绘制马赛克图片
	newImage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))
	// 克隆嵌入图片数据库
	db := cloneTilesDB()
	sp := image.Point{0, 0}
	for y := bounds.Min.Y; y < bounds.Max.Y; y += tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x += tileSize {
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}
			nearest := nearest(color, &db)
			fmt.Println("file: ", nearest)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					// 然后将调整大小后的嵌入图片绘制到当前区块位置，从而真正嵌入到马赛克图片中
					draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("1. error:", err, nearest)
				}
			} else {
				fmt.Println("2. error:", err, nearest)
			}
			file.Close()
		}
	}

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	// 原始图片 base64 的值
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newImage, nil)
	// 马赛克图片 base64 的值
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())

	t1 := time.Now()
	// 构建一个包含原始图片，马赛克图片和处理时间的字典类型变量 images
	images := map[string]string{
		"original": originalStr,
		"mosaic": mosaic,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}

	// 将images值渲染到响应的HTML文件返回给用户
	t, _ := template.ParseFiles("views/results.html")
	t.Execute(w, images)
}

func Upload(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("views/upload.html")
	t.Execute(w, nil)
}
