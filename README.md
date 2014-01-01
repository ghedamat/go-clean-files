# Go clean Files

my first go yakshave

Find all the files under a given path that are older then a given date

warn me via email

delete them

## Usage
```
go get github.com/ghedamat/go-clean-files

go install github.com/ghedamat/go-clean-files

go-clean-files -h
```

## Config File
The script expects a config file
`.go-clean-filesrc` in your home directory
written in <b>json</b> format

```
example:
{
  "server": "smtp.gmail.com",
  "port": 587,
  "username": "mail@gmail.com",
  "password": "password",
  "to": "mail1@gmail.com, mail2@gmail.com",
  "mailThreshold": 90,
  "deleteThreshold": 97,
  "subject": "email subject"
}
```
