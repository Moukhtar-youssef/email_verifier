package utils

import "strings"

func splitDomain(domain string) (string, string) {
	parts := strings.Split(domain, ".")
	n := len(parts)
	if n >= 2 {
		return parts[n-2], parts[n-1]
	}
	return "", parts[0]
}
