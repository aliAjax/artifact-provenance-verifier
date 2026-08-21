package platform

import (
	"sort"
	"strings"
)

func SortComponents(v []Component) []Component {
	out := append([]Component(nil), v...)
	sort.Slice(out, func(i, j int) bool {
		a := strings.ToLower(out[i].Name)
		b := strings.ToLower(out[j].Name)
		if a == b {
			return out[i].Version < out[j].Version
		}
		return a < b
	})
	return out
}
func UniqueStrings(in []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, v := range in {
		if !seen[v] {
			seen[v] = true
			out = append(out, v)
		}
	}
	return out
}
func CountSeverities(v []Vulnerability) map[string]int {
	out := map[string]int{}
	for _, x := range v {
		out[strings.ToLower(x.Severity)]++
	}
	return out
}
func HighestSeverity(v []Vulnerability) string {
	best := ""
	rank := 0
	for _, x := range v {
		if r := SeverityRank(x.Severity); r > rank {
			rank = r
			best = x.Severity
		}
	}
	return best
}
