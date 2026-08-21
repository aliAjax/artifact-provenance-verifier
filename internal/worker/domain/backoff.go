package domain

import "time"

func NextRun(now time.Time, attempt int) time.Time {
	if attempt < 0 {
		attempt = 0
	}
	if attempt > 10 {
		attempt = 10
	}
	return now.Add(time.Duration(1<<attempt) * time.Second)
}
