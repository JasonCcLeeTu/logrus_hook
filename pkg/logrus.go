package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

func NewLogrus(level logrus.Level) {
	l := logrus.New()

	l.SetLevel(level)
	l.SetReportCaller(true)

	formatter := &SysLogFormatter{}

	l.SetFormatter(formatter)

}

type SysLogFormatter struct {
}

func (s *SysLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields)

	for k, v := range entry.Data {
		data[k] = v
	}

	data["file"] = fmt.Sprintf("%v:%v", entry.Caller.File, entry.Caller.Line) // add default file data
	data["function"] = entry.Caller.Function                                  // add default function data

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
