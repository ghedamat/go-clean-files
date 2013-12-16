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

	files := cleaner.GetSortedFiles(root, settings)

	for _, f := range files {
		fmt.Printf("%s %s\n", f.ModTime(), f.Path)
	}
	fmt.Println("files found: ", len(files))

	//cleaner.SendEmail(settings)

}
