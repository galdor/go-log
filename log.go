package log

import (
	"strconv"
	"time"
)

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelError Level = "error"
)

type Message struct {
	Time       *time.Time
	Level      Level
	DebugLevel int
	Message    string
	Data       Data

	domain string
}

func (msg Message) FullLevel() string {
	level := string(msg.Level)

	if msg.Level == LevelDebug {
		level += "." + strconv.Itoa(msg.DebugLevel)
	}

	return level
}

type Datum interface{}

type Data map[string]Datum

func MergeData(dataList ...Data) Data {
	data := Data{}

	for _, d := range dataList {
		for k, v := range d {
			data[k] = v
		}
	}

	return data
}
