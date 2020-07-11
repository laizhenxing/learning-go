package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	deserialization()
}

func deserialization()  {
	f, err := os.Open("./student.xml")
	if err != nil {
		fmt.Println("open student.xml error: ", err)
		return
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stelm := xml.StartElement(tok)
			fmt.Println("start: ", stelm.Name.Local)
		case xml.EndElement:
			endelm := xml.EndElement(tok)
			fmt.Println("end: ", endelm.Name.Local)
		case xml.CharData:
			data := xml.CharData(tok)
			fmt.Println("str: ", string(data))
		}
	}
}
