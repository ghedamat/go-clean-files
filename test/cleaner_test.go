package cleaner_test

import (
	"bytes"
	"github.com/ghedamat/go-clean-files/cleaner"
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"testing"
)

func TestFindFiles(t *testing.T) {
	Convey("Given a directory", t, func() {
		exec.Command("mkdir", "-p", "/tmp/go-clean-files").Run()
		exec.Command("touch", "-d", "2 Jan 2006", "/tmp/go-clean-files/test1").Run()
		exec.Command("touch", "/tmp/go-clean-files/test2").Run()
		path := "/tmp/go-clean-files"
		conf := cleaner.Settings{MailThreshold: 30, DeleteThreshold: 0}
		Convey("when files are searched", func() {
			files := cleaner.GetSortedFiles(path, conf)
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
        "deleteThreshold": 37,
        "mailThreshold": 30
      }`
		rb := new(bytes.Buffer)
		rb.WriteString(file)

		Convey("when the file is parsed", func() {
			conf := cleaner.ReadConf(rb)

			Convey("should have thresholds", func() {
				So(conf.MailThreshold, ShouldEqual, 30)
				So(conf.DeleteThreshold, ShouldEqual, 37)
			})
		})
	})
}
