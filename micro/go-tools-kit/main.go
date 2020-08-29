package main

import (
	"time"

	"go-tools-kit/progressbar"
)

func main() {
	var bar progressbar.Bar
	//bar.NewOption(0, 100)
	start := 31
	total := 100
	bar.NewOptionWithGraph(int64(start), int64(total),"#")
	for i := start; i <= total; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}
