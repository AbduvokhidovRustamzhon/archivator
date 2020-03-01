package main

import (
	"os"
	"time"
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
	ConcurrentArchivator(fileNames)
	SequencedArchivator(fileNames)
	time.Sleep(time.Second)
}