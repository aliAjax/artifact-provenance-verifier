package infrastructure

func WithinLimits(size int, depth int) bool { return size <= 10*1024*1024 && depth <= 64 }
