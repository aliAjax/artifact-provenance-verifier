package adapter

import "strings"

func ValidPURL(v string) bool { return strings.HasPrefix(v, "pkg:") && len(v) > 6 }
