package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	HTMLPagePath string
	Content      string
}

func main() {

	// Creating a flag for user-defined file name
	fileName := flag.String("file", "", "The file to parse")
	flag.Parse()

	// Reading entire file content
	content, err := ioutil.ReadFile(*fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	page := Page{
		// TextFilePath: "/",
		// TextFileName: "first-post.txt",
		HTMLPagePath: *fileName + ".html",
		Content:      string(content),
	}

	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Create a new, blank HTML file.
	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	// Executing the template injects the Page instance's data,
	// allowing us to render the content.
	// Furthermore, upon execution, the rendered template will be
	// saved inside the new file we created earlier.
	t.Execute(newFile, page)

	// Printing page content
	fmt.Println(page.Content)
}
