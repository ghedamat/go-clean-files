package cleaner

import (
	"os"
	"path/filepath"
	"sort"
	"time"
)

// byName implements sort.Interface.
type byDate []MyFile

func (f byDate) Len() int           { return len(f) }
func (f byDate) Less(i, j int) bool { return f[i].Info.ModTime().Before(f[j].Info.ModTime()) }
func (f byDate) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

/*
func getFiles() []os.FileInfo {
	files, _ := ioutil.ReadDir("./test")
	return files
}
*/

func sortFiles(files []MyFile) {
	sort.Sort(byDate(files))
}

type MyFile struct {
	Path string
	Info os.FileInfo
}

func findFiles(path string) []MyFile {
	arr := make([]MyFile, 0)
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			arr = append(arr, MyFile{path, f})
		}
		return nil
	})
	return arr
}

func filterFiles(arr *[]MyFile) {
	files := *arr
	n := sort.Search(len(files), func(i int) bool {
		return files[i].Info.ModTime().After(time.Now().AddDate(0, -3, 0))
	})
	*arr = files[:n]
}

func GetSortedFiles(root string) []MyFile {
	arr := findFiles(root)
	sortFiles(arr)
	filterFiles(&arr)
	return arr
}
