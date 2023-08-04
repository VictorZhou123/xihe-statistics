package utils

import (
	"net"
	"net/url"
	"regexp"
)

func IsValidURL(u string) bool {
	_, err := url.Parse(u)
	return err == nil
}

func IsValidIPAddress(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

func isMatchRegex(pattern string, v string) bool {
	matched, err := regexp.MatchString(pattern, v)
	if err != nil {
		return false
	}

	return matched
}
