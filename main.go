package main

import (
	"flag"
	"fmt"
	"github.com/ghedamat/go-clean-files/cleaner"
)

func main() {
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

	//cleaner.SendEmail(settings, mailFiles)
	//cleaner.DeleteFiles(deleteFiles)

}
