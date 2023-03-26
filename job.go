package job

import "time"

type Job struct {
	Fn   func(...any)
	Args []any
}

func (j Job) run() {
	j.Fn(j.Args...)
}

func (j Job) ScheduleRecurring(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			go j.run()
		}
	}()
}

func (j Job) ScheduleOneTime(at time.Time) {
	delay := time.Until(at)
	timer := time.NewTimer(delay)
	go func() {
		<-timer.C
		go j.run()
		timer.Stop()
	}()
}
