package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type Text struct {
	Paragraph string
}

func main() {
	var flagvar string

	fmt.Print("Enter a file name: ")

	flag.StringVar(&flagvar, "flagvar", "first-post.html", "Entered a file name ")
	flag.Parse()
	fmt.Print(flagvar)
	// Reads content from text file
	FileContents, err := ioutil.ReadFile(flagvar + ".txt")
	if err != nil {
		panic(err)
	}

	// fmt.Print(string(FileContents))

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	fmt.Print(t)
	text := Text{Paragraph: string(FileContents)}
	err = t.Execute(os.Stdout, text)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(flagvar + ".html")

	err = t.Execute(f, text)

	if err != nil {
		panic(err)
	}
}
