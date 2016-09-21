package gourl

import (
	"fmt"
	"net"
	"net/url"
	"regexp"

	"golang.org/x/net/idna"
)

// IsSameHost check is source and url host same host
func IsSameHost(source, loc string, www bool) (bool, error) {
	parsed, err := url.Parse(loc)

	if err != nil {
		return false, err
	}

	if !parsed.IsAbs() {
		return true, nil
	}
	source, err = GetHost(source, www, false)
	if err != nil {
		return false, err
	}
	loc, err = GetHost(loc, www, false)
	if err != nil {
		return false, err
	}

	return loc == source, nil
}

// GetHost get hostname without port from loc. cut www if needed
func GetHost(loc string, www, decode bool) (string, error) {
	parsed, err := url.Parse(loc)
	if err != nil {
		return "", err
	}
	host, _, err := net.SplitHostPort(parsed.Host)
	if err != nil {
		host = parsed.Host
	}
	if www {
		re := regexp.MustCompile(`^www\.`)
		host = re.ReplaceAllString(host, "")
	}

	if decode {
		return idna.ToASCII(host)
	}

	return host, nil
}

// GetProto get protocol from
func GetProto(src string) (string, error) {
	parsed, err := url.Parse(src)
	if err != nil {
		return "", err
	}
	if len(parsed.Scheme) > 0 {
		return parsed.Scheme, nil
	}

	return "", nil
}

// AddHTTP adds http to src if needed. if src has http or https leave it
func AddHTTP(src string) string {
	re := regexp.MustCompile("^https?://")
	if re.MatchString(src) {
		return src
	}

	return fmt.Sprintf("http://%s", src)
}

// ProtoLess remove proto from src
func ProtoLess(src string) string {
	re := regexp.MustCompile(`^(.+://)?(.+)$`)

	return re.ReplaceAllString(src, "$2")
}

// WwwLess remote ww from passed src
func WwwLess(src string) (string, error) {
	proto, err := GetProto(src)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(fmt.Sprintf("^(%s://)?www\\.(.+)$", proto))

	return re.ReplaceAllString(src, "$1$2"), nil

}

// AddWww add www to src url
func AddWww(src string) (string, error) {
	src, err := WwwLess(src)
	if err != nil {
		return "", err
	}

	proto, err := GetProto(src)
	if err != nil {
		return "", err
	}

	// remove proto if exists
	if len(proto) > 0 {
		src = ProtoLess(src)
	}

	host := fmt.Sprintf("www.%s", src)
	if len(proto) > 0 {
		host = fmt.Sprintf("%s://%s", proto, host)
	}

	return host, nil
}
