package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)


type Text struct {
	Paragraph string
}


func main() {

	// Reads content from text file
	FileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Print(string(FileContents))


	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	fmt.Print(t)
	text := Text{Paragraph: string(FileContents)}
	err = t.Execute(os.Stdout,text)

	if err != nil {
		panic(err)
	}

	f, err := os.Create("first-post.html")

	err = t.Execute(f, text)

	if err != nil {
		panic(err)
	}
}
