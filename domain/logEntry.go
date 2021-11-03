package domain

import (
	"log"
	"paulo/parse-log-apache/util/regex"
	"regexp"
	"time"

	"gopkg.in/validator.v2"
)

const patten = `(?P<ip>[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}) - - \[(?P<date>.*)] \"(?P<method>[A-Z]+) (?P<resource>/[0-9a-zA-Z-_?=+,%:~;{}()\[\]&\/\.]*) (?P<version>HTTP/[0-9+.[0-9]+)\" (?P<status>[0-9]{3}) (?P<size>[0-9]+|-) \"(?P<referer>([0-9a-zA-Z-_?=+,%:~;{}()\[\]&\/\.]*|-))\" \"(?P<useragent>.*|-)\" \"`
const datePattern = "02/Jan/2006:15:04:05 +0100"

type LogEntry struct {
	IP        string    `validate:"nonzero"`
	Date      time.Time `validate:"nonzero"`
	Method    string    `validate:"nonzero"`
	Resource  string    `validate:"nonzero"`
	Version   string
	Status    string
	Size      string
	Referer   string
	UserAgent string
}

type LogEntryPostgreSQL struct {
	ID        string `db:"id"`
	IP        string `db:"ip"`
	Date      string `db:"date"`
	Method    string `db:"method"`
	Resource  string `db:"resource"`
	Version   string `db:"version"`
	Status    string `db:"status"`
	Size      string `db:"size"`
	Referer   string `db:"referer"`
	UserAgent string `db:"user_agent"`
}

type LogEntryRepository interface {
	BulkInsert(logs []LogEntry) error
}

func (l LogEntry) Parse(line string) (LogEntry, error) {
	re := regexp.MustCompile(patten)
	match := re.FindStringSubmatch(line)
	groupNames := re.SubexpNames()

	ip := regex.GetValueByGroup("ip", match, groupNames)
	dateString := regex.GetValueByGroup("date", match, groupNames)

	date, err := time.Parse(datePattern, dateString)

	if err != nil {
		log.Println(err)
		return LogEntry{}, err
	}

	method := regex.GetValueByGroup("method", match, groupNames)
	resource := regex.GetValueByGroup("resource", match, groupNames)
	version := regex.GetValueByGroup("version", match, groupNames)
	status := regex.GetValueByGroup("status", match, groupNames)
	size := regex.GetValueByGroup("size", match, groupNames)
	referer := regex.GetValueByGroup("referer", match, groupNames)
	userAgent := regex.GetValueByGroup("useragent", match, groupNames)

	res := LogEntry{
		IP:        ip,
		Method:    method,
		Date:      date,
		Resource:  resource,
		Version:   version,
		Status:    status,
		Size:      size,
		Referer:   referer,
		UserAgent: userAgent,
	}

	if err := validator.Validate(res); err != nil {
		log.Println(err)
		return LogEntry{}, err
	}

	return res, nil
}
