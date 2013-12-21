package cleaner_test

import (
	"bytes"
	"github.com/ghedamat/go-clean-files/cleaner"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestDeleteFiles(t *testing.T) {
	Convey("Given a list of files", t, func() {
		exec.Command("mkdir", "-p", "/tmp/go-clean-files").Run()
		exec.Command("touch", "-d", "2 Jan 2006", "/tmp/go-clean-files/test1").Run()
		exec.Command("touch", "/tmp/go-clean-files/test2").Run()
		exec.Command("touch", "/tmp/go-clean-files/aaa").Run()
		path := "/tmp/go-clean-files"
		files := cleaner.GetSortedFiles(path, 0)
		Convey("when the list is passed", func() {
			err := cleaner.DeleteFiles(files)
			Convey("the files are deleted", func() {
				So(err, ShouldEqual, nil)
				res, _ := ioutil.ReadDir("/tmp/go-clean-files")
				So(res, ShouldResemble, []os.FileInfo{})
			})
		})
	})
}

func TestSendMail(t *testing.T) {
	Convey("HOW?", t, func() {
	})
}

func TestPrepareMessage(t *testing.T) {
	Convey("Given a list of files", t, func() {
		exec.Command("mkdir", "-p", "/tmp/go-clean-files").Run()
		exec.Command("touch", "-d", "2 Jan 2006", "/tmp/go-clean-files/test1").Run()
		exec.Command("touch", "/tmp/go-clean-files/test2").Run()
		exec.Command("touch", "/tmp/go-clean-files/aaa").Run()
		path := "/tmp/go-clean-files"
		files := cleaner.GetSortedFiles(path, 0)

		Convey("when the list is passed", func() {
			msg := cleaner.PrepareMessage(files)
			Convey("a message containing the list is prepared", func() {
				So(msg, ShouldEqual, `/tmp/go-clean-files/aaa
/tmp/go-clean-files/test1
/tmp/go-clean-files/test2`)
			})
		})
	})
}

func TestFindFiles(t *testing.T) {
	Convey("Given a directory", t, func() {
		exec.Command("mkdir", "-p", "/tmp/go-clean-files").Run()
		exec.Command("touch", "-d", "2 Jan 2006", "/tmp/go-clean-files/test1").Run()
		exec.Command("touch", "/tmp/go-clean-files/test2").Run()
		path := "/tmp/go-clean-files"
		Convey("when files are searched", func() {
			files := cleaner.GetSortedFiles(path, 30)
			Convey("files older than X are found", func() {
				So(files[0].Path, ShouldEqual, "/tmp/go-clean-files/test1")
			})
		})
	})
}

func TestConfigParser(t *testing.T) {
	Convey("Given a config file", t, func() {
		file := `
      {
        "server": "smtp.gmail.com",
        "port": 587,
        "username": "email@gmail.com",
        "password": "password",
        "to": "dest1@gmail.com"
      }`
		rb := new(bytes.Buffer)
		rb.WriteString(file)

		Convey("when the file is parsed properly", func() {
			conf := cleaner.ReadConf(rb)

			Convey("all fields are populated", func() {
				So(conf.Server, ShouldEqual, "smtp.gmail.com")
				So(conf.Port, ShouldEqual, 587)
				So(conf.Username, ShouldEqual, "email@gmail.com")
				So(conf.Password, ShouldEqual, "password")
				So(conf.To, ShouldEqual, "dest1@gmail.com")
			})
		})
	})
	Convey("Given a config file with two emails", t, func() {
		file := `
      {
        "to": "dest1@gmail.com,dest2@gmail.com"
      }`
		rb := new(bytes.Buffer)
		rb.WriteString(file)

		Convey("when the file is parsed", func() {
			conf := cleaner.ReadConf(rb)

			Convey("should have a slice with two emails", func() {
				So(conf.ToAddresses(), ShouldResemble, []string{"dest1@gmail.com", "dest2@gmail.com"})
			})
		})
	})
	Convey("Given a config file with two emails separated with comma/spaces", t, func() {
		file := `
      {
        "to": "dest1@gmail.com, dest2@gmail.com"
      }`
		rb := new(bytes.Buffer)
		rb.WriteString(file)

		Convey("when the file is parsed", func() {
			conf := cleaner.ReadConf(rb)

			Convey("should have a slice with two emails", func() {
				So(conf.ToAddresses(), ShouldResemble, []string{"dest1@gmail.com", "dest2@gmail.com"})
			})
		})
	})
	Convey("Given a config file with proper thresholds", t, func() {
		file := `
      {
        "server": "smtp.gmail.com",
        "port": 587,
        "username": "email@gmail.com",
        "password": "password",
        "to": "dest1@gmail.com",
        "deleteThreshold": 30,
        "mailThreshold": 37
      }`
		rb := new(bytes.Buffer)
		rb.WriteString(file)

		Convey("when the file is parsed", func() {
			conf := cleaner.ReadConf(rb)

			Convey("should have thresholds", func() {
				So(conf.MailThreshold, ShouldEqual, 37)
				So(conf.DeleteThreshold, ShouldEqual, 30)
			})
		})
	})
	Convey("Given a config file", t, func() {
		file := `
      {
        "subject" : "test subject"
      }`
		rb := new(bytes.Buffer)
		rb.WriteString(file)

		Convey("when the file is parsed", func() {
			conf := cleaner.ReadConf(rb)

			Convey("should have Subject", func() {
				So(conf.Subject, ShouldEqual, "test subject")
			})
		})
	})
}
