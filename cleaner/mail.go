package cleaner

import (
	"log"
	"net/smtp"
	"strconv"
)

func SendEmail() {
	settings := ReadConf()

	to := "thamayor@gmail.com"
	subject := "subject"
	msg := "message"

	body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + msg

	server := settings.Server + ":" + strconv.Itoa(settings.Port)
	auth := smtp.PlainAuth("", settings.Username, settings.Password, settings.Server)
	err := smtp.SendMail(server, auth, settings.Username, []string{to}, []byte(body))

	if err != nil {
		log.Fatal(err)
	}
}
