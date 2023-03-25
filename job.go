package job

import "time"

type Job struct {
	fn   func(...any)
	args []any
}

func (j Job) Run() {
	j.fn(j.args...)
}

func (j Job) ScheduleRecurring(interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			go j.Run()
		}
	}()
	return ticker
}

func (j Job) ScheduleOneTime(at time.Time) *time.Timer {
	delay := time.Until(at)
	timer := time.NewTimer(delay)
	go func() {
		<-timer.C
		go j.Run()
	}()
	return timer
}