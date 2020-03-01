package main

import (
	"archive/zip"
	"io"
	"os"
	"sync"
)

func SequencedArchivator(fileNames []string) {
	for _, fileName := range fileNames {
		Archive(fileName)
	}
}

func ConcurrentArchivator(fileNames []string) {
	waitGroup := sync.WaitGroup{}
	for _, fileName := range fileNames {
		fName := fileName
		waitGroup.Add(1)
		go func(wg *sync.WaitGroup, fileName string) {
			defer func() {
				waitGroup.Done()
			}()
			Archive(fName)
		}(&waitGroup, fileName)
	}
	waitGroup.Wait()
}

func Archive(fileName string) {
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





