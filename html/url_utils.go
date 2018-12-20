package html

import (
	"net/url"
	"strings"
)

//UriSuffix extracts the suffix from a given Uri, e.g. "suffix" from "https://github.com/a/suffix?b=c#dd
func UriSuffix(url * url.URL) string {
	if url == nil{
		return ""
	}
	path := url.Path
	pathParts := strings.Split(path,"/")
	lastPath := pathParts[len(pathParts)-1]
	return lastPath
}