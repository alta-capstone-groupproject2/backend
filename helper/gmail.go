package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"

	"lami/app/features/users"

	"gopkg.in/gomail.v2"
)

type BodylinkEmail struct {
	SUBJECT string
	Name    string
	URL     string
}

func SendGmailNotify(email, subject string) {
	template, errPath := filepath.Abs("./helper/templates/emailNotif.html")
	fmt.Print(errPath)
	//template := "/home/alfin/ALTA/tugas/capstone/backend/helper/templates/emailNotif.html"

	templateData := BodylinkEmail{
		SUBJECT: subject,
	}
	result, errParse := ParseTemplate(template, templateData)
	fmt.Println(errParse)

	runtime.GOMAXPROCS(1)
	go SendEmail(email, subject, result)
}

func SendEmailVerification(userData users.Core, encrypt string) {
	template, errPath := filepath.Abs("./helper/templates/emailVerify.html")
	fmt.Print(errPath)
	//template := "/home/alfin/ALTA/tugas/capstone/backend/helper/templates/emailVerify.html"
	subject := "Email Verification"

	url := "https://lamiapp.site/users/confirm/" + encrypt

	templateData := BodylinkEmail{
		Name:    userData.Name,
		SUBJECT: subject,
		URL:     url,
	}
	result, errParse := ParseTemplate(template, templateData)
	fmt.Println(errParse)

	runtime.GOMAXPROCS(1)
	go SendEmail(userData.Email, subject, result)
}

func SendEmail(to string, subject string, result string) error {
	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "Lami App <alfin.7007@gmail.com>"
	CONFIG_AUTH_EMAIL := os.Getenv("EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("EMAIL_PASSWORD")
	m := gomail.NewMessage()
	m.SetHeader("From", CONFIG_SENDER_NAME)
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	// m.Attach(templateFile) // attach whatever you want

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
