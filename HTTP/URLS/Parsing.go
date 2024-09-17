package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}
	var parsed ParsedURL
	parsed.protocol = parsedUrl.Scheme
	parsed.username = parsedUrl.User.Username()
	parsed.password,_ = parsedUrl.User.Password()
	parsed.hostname = parsedUrl.Hostname()
	parsed.port = parsedUrl.Port()
	parsed.pathname = parsedUrl.Path
	parsed.search = parsedUrl.RawQuery
	parsed.hash = parsedUrl.Fragment
	return parsed
}
