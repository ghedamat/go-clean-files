package cleaner

import (
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

func SendEmail() {
	settings := ReadConf()

	to := settings.To
	subject := "subject"
	msg := "message"

	body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + msg

	tos := strings.Split(to, ",")

	server := settings.Server + ":" + strconv.Itoa(settings.Port)
	auth := smtp.PlainAuth("", settings.Username, settings.Password, settings.Server)
	err := smtp.SendMail(server, auth, settings.Username, tos, []byte(body))

	if err != nil {
		log.Fatal(err)
	}
}
