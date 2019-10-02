package main

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strconv"
)

func getPhoto(id int) (os.File, error) {
	fileName := strconv.Itoa(id)
	file, err := os.Open("./images/" + fileName + ".jpg")
	if err != nil {
		log.Printf("An error occurred: %v", err)
		return *file, err
	}
	return *file, nil
}


func Download(file multipart.File, id string) (returnErr error) {
	defer func() {
		err := file.Close()

		if err != nil && returnErr == nil {
			log.Printf("error: %v", err)
			returnErr = err
		}
	}()

	tempFile, err := ioutil.TempFile("imagesupload", "upload-*.jpg")
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	defer func() {
		err := tempFile.Close()

		if err != nil && returnErr == nil {
			log.Printf("error: %v", err)
			returnErr = err
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	err = os.Rename(tempFile.Name(), "imagesupload/"+id+".jpg")

	if err != nil {
		log.Printf("An error occurred: %v", err)
		return err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	return nil
}
