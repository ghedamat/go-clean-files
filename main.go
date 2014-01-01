package main

import (
	"flag"
	"fmt"
	"github.com/ghedamat/go-clean-files/cleaner"
)

func main() {
	sendMail := flag.Bool("m", false, "enable email notification")
	remove := flag.Bool("d", false, "enable file removal")
	flag.Parse()
	root := flag.Arg(0)
	settings := cleaner.ParseConf()

	mailFiles := cleaner.GetSortedFiles(root, settings.MailThreshold)
	deleteFiles := cleaner.GetSortedFiles(root, settings.DeleteThreshold)

	for _, f := range mailFiles {
		fmt.Printf("%s %s\n", f.ModTime(), f.Path)
	}
	fmt.Println("mail files found: ", len(mailFiles))
	fmt.Println("delete files found: ", len(deleteFiles))

	if *sendMail {
		fmt.Println("sending Email")
		cleaner.SendEmail(settings, mailFiles)
	}

	if *remove {
		fmt.Println("removing Files")
		cleaner.DeleteFiles(deleteFiles)
	}

}
