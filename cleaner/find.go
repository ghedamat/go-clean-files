package cleaner

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type File struct {
	Path string
	os.FileInfo
}

// byName implements sort.Interface.
type byDate []File

func (f byDate) Len() int           { return len(f) }
func (f byDate) Less(i, j int) bool { return f[i].ModTime().Before(f[j].ModTime()) }
func (f byDate) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func sortFiles(files []File) {
	sort.Sort(byDate(files))
}

func findFiles(path string) []File {
	arr := make([]File, 0)
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !f.IsDir() {
			arr = append(arr, File{path, f})
		}
		return nil
	})
	return arr
}

func filterFiles(arr *[]File, days int) {
	files := *arr
	n := sort.Search(len(files), func(i int) bool {
		return files[i].ModTime().After(time.Now().AddDate(0, 0, -days))
	})
	*arr = files[:n]
}

func GetSortedFiles(root string, conf Settings) []File {
	arr := findFiles(root)
	sortFiles(arr)
	filterFiles(&arr, conf.MailThreshold)
	return arr
}
