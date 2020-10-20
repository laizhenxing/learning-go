package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

var r = rate.NewLimiter(1, 5)

func main() {
	rt := mux.NewRouter()
	{
		rt.Methods("GET").Path("/allown").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("ok"))
		})
	}

	if err := http.ListenAndServe(":8079", AllowNMiddleware(rt)); err != nil {
		fmt.Println("http服务启动出错：", err)
	}
}

func AllowNMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !r.AllowN(time.Now(), 2) {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, req)
	})
}

func allowN()  {
	for {
		if r.AllowN(time.Now(), 2) {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println("too many requests.")
		}
		time.Sleep(time.Second)
	}
}
