package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("should return an error if string is invalid", func(t *testing.T) {
		invalidLine := `86.201.251.107 - - [10/Jan/2021:23:38:12 +0100] "GET /apache-log/access.log'A=0 HTTP/1.1" 231`
		l := new(LogEntry)
		_, err := l.Parse(invalidLine)
		assert.Error(t, err)
	})
	t.Run("should return an error if string is invalid", func(t *testing.T) {
		invalidLine := `13.66.139.0 - - [19/Dec/2020:13:57:26 +0100] "GET /index.php?option=com_phocagallery&view=category&id=1:almhuette-raith&Itemid=53 HTTP/1.1" 200 32653 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)" "-"
		`
		l := new(LogEntry)
		_, err := l.Parse(invalidLine)
		assert.NoError(t, err)
	})
}
