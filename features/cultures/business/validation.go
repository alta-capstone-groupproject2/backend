package business

import (
	"errors"
	"fmt"
	"lami/app/helper"
	"mime/multipart"
	"strconv"
	"time"
)

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
	fmt.Println(fileName)
	// Upload file
	urlImage, errUploadImg := helper.UploadFileToS3(directory, fileName, contentType, fileData)

	if errUploadImg != nil {
		return "", errors.New("failed to upload file")
	}
	fmt.Println(urlImage)
	return urlImage, nil
}
