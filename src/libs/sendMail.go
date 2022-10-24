package libs

import (
	"log"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(receive string, subject string, body interface{}) {
	mail := gomail.NewMessage()

	mail.SetHeader("From", os.Getenv("EMAILUSER"))
	mail.SetHeader("To", receive)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body.(string))

	port_to_int, _ := strconv.Atoi(os.Getenv("EMAILPORT"))

	sending := gomail.NewDialer(os.Getenv("EMAILSERVER"), port_to_int, os.Getenv("EMAILUSER"), os.Getenv("EMAILPASSWORD"))
	if err := sending.DialAndSend(mail); err != nil {
		log.Fatalln(err.Error())
	}
}
