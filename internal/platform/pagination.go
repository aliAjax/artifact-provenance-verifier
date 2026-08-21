package platform

import (
	"fmt"
	"strconv"
	"strings"
)

type Page struct {
	Limit  int
	Cursor string
}

func ParsePage(limitText, cursor string) Page {
	n, _ := strconv.Atoi(limitText)
	if n <= 0 {
		n = 50
	}
	if n > 500 {
		n = 500
	}
	return Page{Limit: n, Cursor: cursor}
}
func EncodeCursor(id string) string { return fmt.Sprintf("c:%s", id) }
func DecodeCursor(v string) (string, bool) {
	if !strings.HasPrefix(v, "c:") {
		return "", false
	}
	x := strings.TrimPrefix(v, "c:")
	return x, x != ""
}
func Paginate[T any](in []T, limit int) ([]T, bool) {
	if limit <= 0 {
		limit = 50
	}
	if len(in) <= limit {
		return in, false
	}
	return in[:limit], true
}
func CursorIndex(ids []string, cursor string) int {
	if cursor == "" {
		return 0
	}
	for i, v := range ids {
		if v == cursor {
			return i + 1
		}
	}
	return 0
}
