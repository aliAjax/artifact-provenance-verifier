package infrastructure

type Allowlist struct{ Subjects map[string]bool }

func (a Allowlist) Allowed(v string) bool { return a.Subjects[v] }
