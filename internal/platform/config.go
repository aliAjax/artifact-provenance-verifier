package platform

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTPAddr       string
	APIKey         string
	DatabaseURL    string
	MaxBody        int64
	RequestTimeout time.Duration
	WorkerInterval time.Duration
}

func LoadConfig() Config {
	c := Config{HTTPAddr: env("HTTP_ADDR", ":8092"), APIKey: env("API_KEY", "dev-secret"), DatabaseURL: env("DATABASE_URL", "memory://"), MaxBody: 4 << 20, RequestTimeout: 10 * time.Second, WorkerInterval: 30 * time.Second}
	if v := os.Getenv("MAX_BODY"); v != "" {
		if n, e := strconv.ParseInt(v, 10, 64); e == nil {
			c.MaxBody = n
		}
	}
	return c
}
func env(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
