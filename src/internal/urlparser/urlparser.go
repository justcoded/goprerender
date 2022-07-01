package urlparser

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"
)

func ParseUrl(queryString string) (string, string, error) {
	u, err := url.Parse(queryString)
	if err != nil {
		return "", "", err
	}

	requestURL := ""
	hostPath := ""

	if u.Path != "/" && strings.HasSuffix(u.Path, "/") {
		path := strings.TrimRight(u.Path, "/")
		requestURL = u.Scheme + "://" + u.Host + path
		hostPath = u.Host + path
	} else {
		requestURL = queryString
		hostPath = u.Host + u.Path
	}

	return requestURL, hostPath, err
}

func GetHashKey(queryString string) (string, error) {
	_, hostPath, err := ParseUrl(queryString)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(hostPath))), err
}
