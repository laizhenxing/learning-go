package main

import (
	"gblog/models"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Cron starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanTag ")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle")
		models.CleanAllArticle()
	})

	c.Start()

	// 会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
