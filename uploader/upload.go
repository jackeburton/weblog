package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"io/ioutil"
	"strings"
	//"html/template"
)

type config_options struct{
	obsidian_home_dir, weblog_tag, properties_separator, property_indicator, title_tag string
}

func main() {
	// this is only going to work if the server is running on my machine or I sync these files to somewhere
	// I should sync the files to somewhere and go from there
	

	config := config_options{
		obsidian_home_dir: "/Users/jackburton/Library/Mobile Documents/iCloud~md~obsidian/Documents/home/Home",
		weblog_tag: "postable-weblog-entry",
		properties_separator: "---",
		property_indicator: "tags",
		title_tag: "Title:",
	}
	
	files_to_upload := get_files_to_upload(config)
	fmt.Print("Uploading files : ")
	for _, file := range files_to_upload {
		fmt.Printf("%+v\n", file)
	}
	//refreshes the aws s3 instance where my blog files are hosted
	// upload_files()
}

func get_files_to_upload(config config_options) []os.DirEntry{
	entry_files := []os.DirEntry{} 
	files, err := os.ReadDir(config.obsidian_home_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		// do this recursively so we can have directories -> filepath.Walk
		// allow for the option of chosing which directory to use as the blog path or just take the home dir?
		if file.IsDir(){
			continue
		}
		file_contents := get_file_contents(config.obsidian_home_dir, file)

		if !strings.Contains(file_contents, config.property_indicator) {
			continue
		}
		// get the tags from the file
		split_contents := strings.Split(file_contents, config.properties_separator)
		properties := split_contents[1]
		
		if !strings.Contains(properties, config.weblog_tag){
			continue
		}

		if !strings.Contains(properties, config.title_tag){
			fmt.Print("file : " + file.Name() + " does not contain a title")
			continue
		}

		// at this point we can upload the file and once uploaded we trigger the serve section
		// so lets put them into another slice to then loop through when we want to serve them
		entry_files = append(entry_files, file)

		/*
		entry_body := split_contents[2]
		entry_title := get_entry_title(properties, title_tag)
		
		fmt.Print(weblog_title)
		fmt.Print(weblog_body)
		new_entry := entry{title: weblog_title, content: weblog_body}
		entries = entries.append()
		*/
	} 
	return entry_files

}

func get_file_contents(dir string, file os.DirEntry) string{

	file_path := filepath.Join(dir, file.Name())
	file_contents, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file_contents.Close()

	// speed improvement -> read only a certain number of bytes to get the tag and separate this to new func
	content, err := ioutil.ReadAll(file_contents)
	if err != nil {
		log.Fatal(err)
	}
	
	return string(content)
}

func get_entry_title(properties string, title_tag string) string {
	return strings.Split(strings.Split(properties, title_tag + " ")[1], `"`)[0]
}
