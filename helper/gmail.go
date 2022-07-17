package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"runtime"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Lami App <alfin.7007@gmail.com>"

type BodylinkEmail struct {
	Name string
	URL string
}

func Send(gmail string, code string) error{
	
	templateData := BodylinkEmail{
		URL: "https://detik.id/",
	}

	to := gmail
	runtime.GOMAXPROCS(1)
	go SendEmailVerification(to, templateData)

	return nil
}

func SendEmailVerification(to string, data interface{}) {
	template := "/home/alfin/ALTA/tugas/capstone/backend/helper/templates/emailVerify.html"
	subject := "Email Verification"

	result, _ := ParseTemplate(template, data)

	err := SendEmail(to, subject, data, result)
	if err == nil {
		fmt.Println("send email '" + subject + "' success")
	} else {
		fmt.Println(err)
	}
}

func SendEmail(to string, subject string, data interface{}, result string) error {
	
	m := gomail.NewMessage()
	m.SetHeader("From", CONFIG_SENDER_NAME)
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	// m.Attach(templateFile) // attach whatever you want


	CONFIG_AUTH_EMAIL := os.Getenv("EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("EMAIL_PASSWORD")
	d := gomail.NewDialer(
		CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}


