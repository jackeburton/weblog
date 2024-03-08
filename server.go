package main

import (
	"os"
	"log"
	"bufio"
	"html/template"
	"net/http"
	"strings"
	"fmt"
)



func StartServer() {

	config := config_options{
		obsidian_home_dir: "/Users/jackburton/Library/Mobile Documents/iCloud~md~obsidian/Documents/home/Home",
		weblog_tag: "postable-weblog-entry",
		properties_separator: "---",
		property_indicator: "tags",
		title_tag: "Title:",
		template: "templates/template.html",
		markdown_dir : "markdown", 
	}

	// get the markdown files from the markdown folder
	files, err := os.ReadDir(config.markdown_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range(files){
		title, content := get_markdown_content(file, config)
		// convert the markdown files into html
		var html_template, err = template.ParseFiles(config.template)
		if err != nil {
			log.Fatal(err)
		}
		// serve the markdown files
		handler := func(w http.ResponseWriter, r *http.Request){
			html_template.Execute(w, content)
		}
		http.HandleFunc("/" + title, handler)
		fmt.Println(title)
	}
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
	fmt.Println("running")
}

func get_markdown_content(file os.DirEntry, config config_options)(string, string){
	md_file, err := os.Open(config.markdown_dir + "/" + file.Name())
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(md_file)
	

	var line string
	var property_buffer []string
	var content_buffer []string
	var title string
	var content string
	separator_count := 0
	for scanner.Scan(){
		line = scanner.Text()
		if len(line) == 0 {
			// paragraph
			continue
		}
		// find title indicator under property indicator -> title of entry [NEED]
		if line == config.properties_separator{
			// everything before should be looped through to find title / tags
			separator_count++ 	
			fmt.Println(separator_count)
			continue
		}
		// find content [NEED]
		// process content, change links to server links [WANT]

		if separator_count < 2 {
			property_buffer = append(property_buffer, line)	
		} else {
			content_buffer = append(content_buffer, line)
		}
	}

	fmt.Println("finished")
	fmt.Println(len(property_buffer))
	fmt.Println(len(content_buffer))

	for _, property_line := range property_buffer{
		if strings.Contains(property_line, config.title_tag){
			// REFACTOR, we already do this in uploader
			fmt.Println(strings.Split(property_line, config.title_tag + " "))
			title = strings.Split(strings.Split(property_line, config.title_tag + " ")[1], `"`)[0]
		}
		// find the tags -> remove postable entry identifier, add rest as tags [WANT]
		// ...
	}

	for _, content_line := range content_buffer{
		content = content + content_line
	}
	return title, content 
}

