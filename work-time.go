package main

import "time"

type WorkTime struct {
	IsWorkTime bool
	Start      time.Time
	End        time.Time
}

func (w *WorkTime) CheckTimePeriod(check time.Time) bool {
	if w.Start.Before(w.End) {
		return !check.Before(w.Start) && !check.After(w.End)
	}
	if w.Start.Equal(w.End) {
		return check.Equal(w.Start)
	}
	return !w.Start.After(check) || !w.End.Before(check)
}
