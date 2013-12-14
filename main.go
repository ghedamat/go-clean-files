package main

import (
	"flag"
	"fmt"
	"github.com/ghedamat/go-clean-files/cleaner"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	files := cleaner.GetSortedFiles(root)

	for _, f := range files {
		fmt.Printf("%s %s\n", f.Info.ModTime(), f.Path)
	}
	fmt.Println("files found: ", len(files))

	cleaner.SendEmail()

}
