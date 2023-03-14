package common

import (
	"bytes"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type MyFormatter struct {
}

func (m *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	newLog := fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}
