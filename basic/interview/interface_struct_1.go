package main

import "fmt"

type Peo interface {
	Speak(string) string
}

type Stu struct {
}

func (s *Stu) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// 结构体没有实现 Peo 这个接口，只有结构体指针实现了
	var p Peo = Stu{}	// cannot use Stu literal (type Stu) as type Peo in assignment: Stu does not implement Peo (Speak method has pointer receiver)
	think := "speak"
	fmt.Println(p.Speak(think))
}