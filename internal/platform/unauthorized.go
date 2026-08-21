package platform

func IsUnauthorized(e error) bool { return e != nil && e.Error() == ErrUnauthorized.Error() }
