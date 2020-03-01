package main

import (
	"os"
)

const (
	zipFormat    = ".zip"
	filesPath    = "./files/"
	archivesPath = "./archives/"
)

func main() {
	fileNames := os.Args[1:]
	if fileNames == nil {
		return
	}
	ConArchivator(fileNames)
	SeqArchivator(fileNames)
	//time.Sleep(time.Second)
}