package platform

func IsNotFound(e error) bool { return e != nil && e.Error() == ErrNotFound.Error() }
