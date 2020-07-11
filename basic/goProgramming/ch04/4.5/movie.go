package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`	// Tag 第一部分的值用于制定JSON对象的名字
	Color  bool `json:"color, omitempty"`	// Tag omitempty 表示结构体成员为空或者零值时不生成JSON对象
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	marshalIndex()
}

func marsha()  {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON Marshal failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func marshalIndex() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalIndex failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
