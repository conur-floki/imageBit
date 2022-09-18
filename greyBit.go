package main

import "os"

func main() {
	if len(os.Args) < 2 {
		log.fatalln("Image path is required")
	}
}
