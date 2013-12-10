package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// byName implements sort.Interface.
type byDate []os.FileInfo

func (f byDate) Len() int           { return len(f) }
func (f byDate) Less(i, j int) bool { return f[i].ModTime().Before(f[j].ModTime()) }
func (f byDate) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func getFiles() []os.FileInfo {
	files, _ := ioutil.ReadDir("./test")
	return files
}

func sortFiles(files []os.FileInfo) {
	sort.Sort(byDate(files))
}

func main() {
	files := getFiles()
	for _, f := range files[0:10] {
		fmt.Println(f.ModTime(), f.Name())
	}

}
