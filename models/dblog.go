package models

import "github.com/apex/log"

type DBLog struct {
}

func (l *DBLog) Printf(format string, v ...interface{}) {
	log.Debugf(format, v)
}
