package platform

func IsConflict(e error) bool { return e != nil && e.Error() == ErrConflict.Error() }
