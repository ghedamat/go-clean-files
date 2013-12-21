package cleaner

import (
	"log"
	"net/smtp"
	"sort"
	"strconv"
	"strings"
)

// byName implements sort.Interface.
type byName []File

func (f byName) Len() int           { return len(f) }
func (f byName) Less(i, j int) bool { return f[i].Name() < f[j].Name() }
func (f byName) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func sortFilesByName(files []File) {
	sort.Sort(byName(files))
}

func PrepareMessage(files []File) string {
	sortFilesByName(files)
	names := make([]string, 0)
	for _, f := range files {
		names = append(names, f.Path)
	}
	return strings.Join(names, "\n")
}

func SendEmail(settings Settings, files []File) {
	to := settings.To

	subject := ""
	if settings.Subject != "" {
		subject = settings.Subject
	} else {
		subject = "File Warning"
	}
	msg := PrepareMessage(files)

	body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + msg

	tos := settings.ToAddresses()

	server := settings.Server + ":" + strconv.Itoa(settings.Port)
	auth := smtp.PlainAuth("", settings.Username, settings.Password, settings.Server)
	err := smtp.SendMail(server, auth, settings.Username, tos, []byte(body))

	if err != nil {
		log.Fatal(err)
	}
}
