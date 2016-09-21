package gourl

import "testing"

type testpair struct {
	src string
	url string
}

var testsSame = []testpair{
	{"http://mail.ru:8080", "https://mail.ru"},
	{"http://mail.ru", "https://www.mail.ru"},
	{"http://mail.ru", "mail.ru"},
	{"www.mail.ru", "mail.ru"},
	{"lo.ru", "ya.ru"},
}

var testsDiff = []testpair{
	{"http://lo.ru:8080", "http://ya.ru:80"},
}

var testsHost = []testpair{
	{"http://mail.ru:70", "mail.ru"},
	{"http://mail.ru:80/fasdfasdfasdf.com", "mail.ru"},
}

var testsHostFail = []testpair{
	{"lo.ru", "/fsdfadsfa"},
	{"new-sebastopol.com", "mail.ru"},
	{"http://new-sebastopol.com", "mail.ru"},
	{"https://new-sebastopol.com", "http://new-sebastopol.com"},
}

var testsProto = []testpair{
	{"http://mail.ru", "http"},
	{"hmail.ru/fasdfasdfasdf.com", ""},
	{"ftp://mail.ru/fasdfasdfasdf.com", "ftp"},
	{"https://hmail.ru/fasdfasdfasdf.com", "https"},
}

var testsAddHTTP = []testpair{
	{"http://mail.ru", "http://mail.ru"},
	{"https://mail.ru", "https://mail.ru"},
	{"mail.ru/fasdfasdfasdf.com", "http://mail.ru/fasdfasdfasdf.com"},
}

var testsWwwLess = []testpair{
	{"http://www.mail.ru", "http://mail.ru"},
	{"https://mail.ru", "https://mail.ru"},
	{"https://www.mail.ru", "https://mail.ru"},
	{"www.mail.ru/fasdfasdfasdf.com", "mail.ru/fasdfasdfasdf.com"},
	{"ftp://www.mail.ru/fasdfasdfasdf.com", "ftp://mail.ru/fasdfasdfasdf.com"},
	{"ftp://mail.ru/fasdfasdfasdf.com", "ftp://mail.ru/fasdfasdfasdf.com"},
}

var testsAddWww = []testpair{
	{"http://www.mail.ru", "http://www.mail.ru"},
	{"http://mail.ru", "http://www.mail.ru"},
	{"https://mail.ru", "https://www.mail.ru"},
	{"https://www.mail.ru", "https://www.mail.ru"},
	{"www.mail.ru/fasdfasdfasdf.com", "www.mail.ru/fasdfasdfasdf.com"},
	{"ftp://www.mail.ru/fasdfasdfasdf.com", "ftp://www.mail.ru/fasdfasdfasdf.com"},
	{"ftp://mail.ru/fasdfasdfasdf.com", "ftp://www.mail.ru/fasdfasdfasdf.com"},
}

var testsProtoLess = []testpair{
	{"http://www.mail.ru", "www.mail.ru"},
	{"https://mail.ru", "mail.ru"},
	{"ftp://www.mail.ru", "www.mail.ru"},
	{"www.mail.ru/fasdfasdfasdf.com", "www.mail.ru/fasdfasdfasdf.com"},
	{"ftp://www.mail.ru/fasdfasdfasdf.com", "www.mail.ru/fasdfasdfasdf.com"},
}

func TestDiffHostT(t *testing.T) {
	for _, pair := range testsDiff {
		if ok, _ := IsSameHost(pair.src, pair.url, true); ok {
			t.Errorf("fail diff %s -> %s", pair.src, pair.url)
		}
	}
}

func TestSameHost(t *testing.T) {
	for _, pair := range testsSame {
		if ok, _ := IsSameHost(pair.src, pair.url, true); !ok {
			t.Errorf("fail %s -> %s", pair.src, pair.url)
		}
	}
}

func TestDiffHost(t *testing.T) {
	for _, pair := range testsDiff {
		if ok, _ := IsSameHost(pair.src, pair.url, true); ok {
			t.Errorf("fail diff %s -> %s", pair.src, pair.url)
		}
	}
}

func TestGetHost(t *testing.T) {
	for _, pair := range testsHost {
		if host, _ := GetHost(pair.src, true, false); host != pair.url {
			t.Errorf("fail get host %s -> %s ,  exptected %s", pair.src, host, pair.url)
		}
	}
}

func TestGetHostFail(t *testing.T) {
	for _, pair := range testsHostFail {
		if host, _ := GetHost(pair.src, true, false); host == pair.url {
			t.Errorf("fail get host %s -> %s ,  exptected %s", pair.src, host, pair.url)
		}
	}
}

func TestGetProto(t *testing.T) {
	for _, pair := range testsProto {
		if proto, err := GetProto(pair.src); err != nil || proto != pair.url {
			t.Errorf("fail get proto %s -> %s (%s) ,  exptected %s", pair.src, proto, err, pair.url)
		}
	}
}

func TestAddHTTP(t *testing.T) {
	for _, pair := range testsAddHTTP {
		if host := AddHTTP(pair.src); host != pair.url {
			t.Errorf("fail add http %s -> %s ,  exptected %s", pair.src, host, pair.url)
		}
	}
}

func TestWwwLess(t *testing.T) {
	for _, pair := range testsWwwLess {
		if res, _ := WwwLess(pair.src); res != pair.url {
			t.Errorf("fail cut www %s -> %s ,  exptected %s", pair.src, res, pair.url)
		}
	}
}

func TestAddWww(t *testing.T) {
	for _, pair := range testsAddWww {
		if host, _ := AddWww(pair.src); host != pair.url {
			t.Errorf("fail add www %s -> %s ,  exptected %s", pair.src, host, pair.url)
		}
	}
}

func TestProtoLess(t *testing.T) {
	for _, pair := range testsProtoLess {
		if res := ProtoLess(pair.src); res != pair.url {
			t.Errorf("fail cut proto %s -> %s ,  exptected %s", pair.src, res, pair.url)
		}
	}
}
