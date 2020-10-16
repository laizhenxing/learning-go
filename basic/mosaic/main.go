package main

import (
	"fmt"
	"net/http"

	"mosaic/sync"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/upload", sync.Upload)
	mux.HandleFunc("/mosaic", sync.Mosaic)
	server := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	// 初始化图片数据库
	sync.TILESDB = sync.TilesDB()
	fmt.Println("DB: ", sync.TILESDB)
	fmt.Println("图片马赛克应用服务器已经启动...")
	server.ListenAndServe()
}
