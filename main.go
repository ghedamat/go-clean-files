package main

import (
	"flag"
	"fmt"
	"github.com/ghedamat/go-clean-files/cleaner"
	"os"
)

func main() {
	sendMail := flag.Bool("m", false, "enable email notification")
	remove := flag.Bool("d", false, "enable file removal")
	verbose := flag.Bool("v", false, "show files that will be deleted")
	flag.Parse()
	root := flag.Arg(0)

	if root == "" {
		fmt.Println("usage: go-clean-files [-v|-d|-m] PATH")
		os.Exit(0)
	}
	settings := cleaner.ParseConf()

	mailFiles := cleaner.GetSortedFiles(root, settings.MailThreshold)
	deleteFiles := cleaner.GetSortedFiles(root, settings.DeleteThreshold)

	if *verbose {
		for _, f := range mailFiles {
			fmt.Printf("%s %s\n", f.ModTime(), f.Path)
		}
	}

	if *sendMail {
		fmt.Println("sending Email")
		cleaner.SendEmail(settings, mailFiles)
	}

	if *remove {
		fmt.Println("removing Files")
		cleaner.DeleteFiles(deleteFiles)
	}

	fmt.Println("mail files found: ", len(mailFiles))
	fmt.Println("delete files found: ", len(deleteFiles))

}
