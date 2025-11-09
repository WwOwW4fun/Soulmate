package utils

import "net/url"

func Escape(s string) string {
	return url.QueryEscape(s)
}
