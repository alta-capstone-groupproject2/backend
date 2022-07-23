package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const UserImages = "userimages"
const UserDocuments = "userdocuments"
const EventImages = "eventimages"
const EventDocuments = "eventdocuments"
const AttendeesDocuments = "attendeesDocuments"
const ProductImages = "productimages"
const CultureImages = "cultureimages"
const ContentImage = "images"
const ContentDocuments = "application/pdf"

var theSession *session.Session

//GetConfig Initiatilize config in singleton way
func GetSession() *session.Session {

	if theSession == nil {
		theSession = initSession()
	}

	return theSession
}

func initSession() *session.Session {
	newSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""),
	}))
	return newSession
}
