package business

import (
	"errors"
	"lami/app/helper"
	"mime/multipart"
	"regexp"
	"strconv"
	"time"
)

func emailFormatValidation(email string) error {
	//	Check syntax email address
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(email))
	if !matched {
		return errors.New("failed syntax email address")
	}
	return nil
}

func phoneFormatValidation(email string) error {
	//	Check syntax email address
	pattern := `^[0-9]+$`
	matched, _ := regexp.Match(pattern, []byte(email))
	if !matched {
		return errors.New("failed syntax phone ")
	}
	return nil
}

func nameFormatValidation(name string) error {
	//	Check syntax email address
	pattern := `^[a-zA-Z ]+$`
	matched, _ := regexp.Match(pattern, []byte(name))
	if !matched {
		return errors.New("failed syntax name")
	}
	return nil
}

func cityFormatValidation(name string) error {
	//	Check syntax email address
	pattern := `^[a-zA-Z ]+$`
	matched, _ := regexp.Match(pattern, []byte(name))
	if !matched {
		return errors.New("failed syntax name")
	}
	return nil
}

func uploadFileValidation(name string, id int, directory string, contentType string, fileInfo *multipart.FileHeader, fileData multipart.File) (string, error) {
	//	Check file extension
	extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename, contentType)
	if err_check_extension != nil {
		return "", errors.New("file extension error")
	}

	//	Check file size
	err_check_size := helper.CheckFileSize(fileInfo.Size, contentType)
	if err_check_size != nil {
		return "", errors.New("file size error")
	}

	//	Memberikan nama file
	fileName := strconv.Itoa(id) + "_" + name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

	// Upload file
	urlImage, errUploadImg := helper.UploadFileToS3(directory, fileName, contentType, fileData)

	if errUploadImg != nil {
		return "", errors.New("failed to upload file")
	}
	return urlImage, nil
}
