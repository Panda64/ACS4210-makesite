package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	HTMLPagePath string
	Content      string
}

func main() {

	// --------------------------------------------------------------------
	// Flags
	// --------------------------------------------------------------------

	// Creating a flag for user-defined file name
	fileName := flag.String("file", "", "The file to parse")

	// Creating a flag for user-defined input directory
	inputDir := flag.String("dir", "", "The input directory")

	flag.Parse()

	// --------------------------------------------------------------------
	// Main
	// --------------------------------------------------------------------

	if *fileName != "" {
		create_html(*fileName)
	} else if *inputDir != "" {
		files, err := ioutil.ReadDir(*inputDir)
		if err != nil {
			ct.Foreground(ct.Red, false)
			log.Fatal(err)
		}

		for _, file := range files {
			if filepath.Ext(file.Name()) == ".txt" {
				create_html(strings.TrimSuffix(file.Name(), ".txt"))
			}
		}
	} else {
		ct.Foreground(ct.Red, false)
		log.Fatal("Please specify a file or directory!")
	}

	ct.Foreground(ct.Green, false)
	fmt.Println("Site Generation Complete!")
	ct.ResetColor()
}

func create_html(fileName string) {
	// Reading entire file content
	content, err := ioutil.ReadFile(fileName + ".txt")
	if err != nil {
		ct.Foreground(ct.Red, false)
		log.Fatal(err)
	}

	page := Page{
		// TextFilePath: "/",
		// TextFileName: "first-post.txt",
		HTMLPagePath: fileName + ".html",
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
	ct.Foreground(ct.Blue, false)
	fmt.Println(page.Content)
}
