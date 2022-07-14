package business

import (
	"errors"
	"lami/app/config"
	"lami/app/features/users"
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

func uploadFileValidation(userReq users.Core, id int, fileInfo *multipart.FileHeader, fileData multipart.File) (string, error) {
	//	Check file extension
	extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return "", errors.New("file extension error")
	}

	//	Check file size
	err_check_size := helper.CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return "", errors.New("file size error")
	}

	//	Memberikan nama file
	fileName := strconv.Itoa(id) + "_" + userReq.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

	// Upload file
	urlImage, errUploadImg := helper.UploadFileToS3(config.UserImages, fileName, fileData)

	if errUploadImg != nil {
		return "", errors.New("failed to upload file")
	}
	return urlImage, nil
}
