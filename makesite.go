package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type Text struct {
	Paragraph string
}

func generateMarkdown(html string) {
	converter := md.NewConverter("", true, nil)

	html = `<strong>Important</strong>`

	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	os.Create(markdown + ".md")

}

func generateHTML(contents []byte, name string) {
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	fmt.Print(t)
	text := Text{Paragraph: string(contents)}
	err := t.Execute(os.Stdout, text)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(name + ".html")

	err = t.Execute(f, text)

	if err != nil {
		panic(err)
	}
	generateMarkdown(name)
}

func main() {
	var flagvar string
	var dir string

	fmt.Print("Enter a file name: ")

	flag.StringVar(&dir, "dir", ".", "Entered a directory ")

	files, _ := ioutil.ReadDir(dir)

	flag.StringVar(&flagvar, "flagvar", "first-post.html", "Entered a file name ")
	flag.Parse()
	fmt.Print(flagvar)
	// Reads content from text file

	for _, s := range files {
		extention := s.Name()[len(s.Name())-4:]
		name := s.Name()[:len(s.Name())-4]
		if extention == ".txt" {
			FileContents, err := ioutil.ReadFile(s.Name())
			if err != nil {
				panic(err)
			}
			generateHTML(FileContents, name)
		}
	}

}
