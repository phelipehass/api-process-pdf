package models

import (
	"4d63.com/tz"
	"strings"
	"time"
)

type ProcessData struct {
	TotalPages  int                `json:"paginas,omitempty"`
	TotalResult int                `json:"total,omitempty"`
	ActualPage  int                `json:"paginaAtual,omitempty"`
	Diaries     []DiaryUnprocessed `json:"diarios,omitempty"`
}

type DiaryUnprocessed struct {
	Diary int64         `json:"diario"`
	Ata   int64         `json:"ata"`
	Date  DiaryDateTime `json:"data"`
}

type DiaryDateTime time.Time

// UnmarshalJSON date format from the diaries
func (dt *DiaryDateTime) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return
	}

	zone, _ := tz.LoadLocation("America/Sao_Paulo")
	s := strings.Trim(string(b), `"`)
	nt, err := time.ParseInLocation("02/01/2006", s, zone)
	*dt = DiaryDateTime(nt)

	return
}

// MarshalJSON date format from the diaries
func (dt *DiaryDateTime) MarshalJSON() (b []byte, err error) {
	if dt == nil {
		return
	}

	timeParse := time.Time(*dt)

	return []byte(timeParse.Format(`"02/01/2006"`)), nil
}
