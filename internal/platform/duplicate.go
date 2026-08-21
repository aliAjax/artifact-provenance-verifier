package platform

func IsDuplicate(e error) bool { return e != nil && e.Error() == ErrDuplicate.Error() }
