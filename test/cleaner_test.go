package cleaner_test

import (
	"bytes"
	"github.com/ghedamat/go-clean-files/cleaner"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

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
}