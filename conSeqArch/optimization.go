package main

import (
	"archive/zip"
	"io"
	"os"
	"sync"
)

func SeqArchivator(fileNames []string) {
	for _, fileName := range fileNames {
		Archivator(fileName)
	}
}

func ConArchivator(fileNames []string) {
	waitGroup := sync.WaitGroup{}
	for _, fileName := range fileNames {
		filesName := fileName
		waitGroup.Add(1)
		go func(fileName string) {
			defer waitGroup.Done()
			Archivator(filesName)
		}(fileName)
	}
	waitGroup.Wait()
}

func Archivator(fileName string) {
	archiveFile, err := os.Create(archivesPath + fileName + zipFormat)
	if err != nil {
		return
	}

	defer func() {
		err = archiveFile.Close()
		if err != nil {
			return
		}
	}()

	writer := zip.NewWriter(archiveFile)

	defer func() {
		err = writer.Close()
		if err != nil {
			return
		}
	}()

	file, err := os.Open(filesPath + fileName)
	if err != nil {
		return
	}

	archive, err := writer.Create(fileName)
	if err != nil {
		return
	}
	_, err = io.Copy(archive, file)
	if err != nil {
		return
	}
}
