package infrastructure

import "time"

func RetryDelay(n int) time.Duration {
	if n < 0 {
		n = 0
	}
	if n > 8 {
		n = 8
	}
	return time.Duration(1<<n) * time.Second
}
