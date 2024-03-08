package main

import (
	"os"
	"fmt"
)

type config_options struct{
	obsidian_home_dir, weblog_tag, properties_separator, property_indicator, title_tag, template, markdown_dir string
}

func main() {
	arg := os.Args[1]
	if arg == "upload"{
		fmt.Println("Uploading files...")
		StartUploader()
		fmt.Println("Files uploaded...")
	} else if arg == "serve"{
		fmt.Println("Starting server...")
		StartServer()
	} else {
		fmt.Println("Command line argument not recognised...")
	}
}
