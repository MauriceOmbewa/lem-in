package main

import (
	"lem-in/root"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Error: usage go run . <input_file(.txt)>")
		return
	}

	root.PrintOutPut()
}
