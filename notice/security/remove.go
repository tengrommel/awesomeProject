package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("/etc/passwd")
	check(err)
	defer file.Close()

	//data, err := ioutil.ReadAll(file)
	//check(err)
	//fmt.Println(string(data))
	//
	//err = ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	//check(err)
	//err = os.Remove("test.txt")
	//check(err)

	//outFile, err := os.Create("test.zip")
	//defer outFile.Close()
	//check(err)
	//zipWriter := zip.NewWriter(outFile)
	//defer zipWriter.Close()
	//var filesToArchive = []struct{
	//	Name, Body string
	//}{
	//{"test.txt", "This is my zipped file!"},
	//{"test2.txt", "This is my zipped file two!"},
	//}
	//for _, file := range filesToArchive {
	//	fileWriter, err := zipWriter.Create(file.Name)
	//	check(err)
	//	_, err = fileWriter.Write([]byte(file.Body))
	//	check(err)
	//}

	zipReader, err := zip.OpenReader("test.zip")
	check(err)
	defer zipReader.Close()

	for _, file := range zipReader.Reader.File {
		zippedFile, err := file.Open()
		defer zippedFile.Close()
		check(err)
		//fileData, _ := ioutil.ReadAll(zippedFile)
		//fmt.Println(string(fileData))
		targetDirectory := "./"
		extractedFilePath := filepath.Join(targetDirectory, file.Name)
		if file.FileInfo().IsDir() {
			log.Println(extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("extracting file:", file.Name)
			outputFile, err := os.OpenFile(extractedFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			check(err)
			defer outputFile.Close()
			_, err = io.Copy(outputFile, zippedFile)
			check(err)
		}
	}
}
