package iptax

import (
	"fmt"
	"time"
)

const (
	PeriodLayout = "2006-01-02"
	MonthEnd     = 25
)

type Period struct {
	From time.Time
	To   time.Time
}

func (p Period) String() string {
	return fmt.Sprintf("%s %d", p.From.Month().String(), p.From.Year())
}

func NewPeriod(now string, periodEnd int) (Period, error) {
	nowdt, err := time.Parse(PeriodLayout, now)
	if err != nil {
		return Period{}, err
	}
	fromdt := time.Date(nowdt.Year(), nowdt.Month()-1, periodEnd,
		0, 0, 0, 0, time.Local)
	todt := time.Date(nowdt.Year(), nowdt.Month(), periodEnd,
		23, 59, 59, 99999, time.Local)
	return Period{From: fromdt, To: todt}, nil
}
