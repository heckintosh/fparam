package utils

import (
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type RequestPrep struct {
	Url    string
	Method string
	Header http.Header
}

func Stable_request(_url string, header http.Header) string {
	var schemes []string
	var resp *http.Response
	var err2 error
	u, err := url.Parse(_url)
	if err != nil {
		log.Error(err)
	}
	scheme := u.Scheme
	//var redirects_allowed bool = false // Support config later on
	if scheme == "https" {
		schemes = []string{"https", "http"}
	} else {
		schemes = []string{"http", "https"}
	}
	client := http.Client{Timeout: 10 * time.Second}
	for _, scheme := range schemes {
		req, err := http.NewRequest("GET", scheme+"://"+u.Host+u.Path, nil)
		req.Header = header
		if err != nil {
			log.Error("Stable request failed: %v", err)
		} else {
			resp, err2 = client.Do(req)
		}
		if err2 != nil {
			var tmp interface{} = err2
			if _, ok := tmp.(net.Error); ok {
				log.Error("Connection error %v:, err")
			} else {
				continue // If not connection error then continue looping in schemes
			}
		} else {
			return resp.Request.URL.String()
		}
		return ""
	}
	return ""
}

func Prepare_request(urls []string) []RequestPrep {
	var reqs []RequestPrep
	header := http.Header{
		"User-Agent":                {"Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0"},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"Accept-Language":           {"en-US,en;q=0.5"},
		"Accept-Encoding":           {"gzip, deflate"},
		"Connection":                {"close"},
		"Upgrade-Insecure-Requests": {"1"},
	}
	methods := []string{"GET", "POST", "JSON"}
	for _, method := range methods {
		for _, url := range urls {
			req := RequestPrep{
				Url:    url,
				Method: method,
				Header: header,
			}
			reqs = append(reqs, req)
		}
	}
	return reqs
}

func Random_str(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Remove_tags(html string) string {
	r := regexp.MustCompile("(?s)<.*?>")
	replaced := r.ReplaceAllString(html, "")
	return replaced
}

func GetKeys(samplemap map[string][]string) []string {
	keys := []string{}
	for key, _ := range samplemap {
		keys = append(keys, key)
	}

	return keys
}

func IsEqual(a1 []string, a2 []string) bool {
	sort.Strings(a1)
	sort.Strings(a2)
	if len(a1) == len(a2) {
		for i, v := range a1 {
			if v != a2[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func extract_headers(headers string) map[string]string {
	tmp := strings.ReplaceAll(headers, "\\n", "\n")
	return parse_headers(tmp)
}
func parse_headers(header string) map[string]string {
	//Remove spaces,tab and return [] string
	res := strings.Fields(header)

	result := make(map[string]string)
	for i := range res {
		if (len(res[i])) > 1 {
			splitted := strings.Split(res[i], ":")
			result[splitted[0]] = splitted[1]
		}
	}
	return result
}

func populate(array []string) map[string]string {
	//converts a list of params into param and value pair
	result := make(map[string]string)
	for i := range array {
		result[array[i]] = strings.Repeat("1", 6-len(strconv.Itoa(i))) + strconv.Itoa(i)
	}
	return result
}
func diff_map(body_1 string, body_2 string) []string {
	var sig []string
	line_1 := strings.Split(body_1, "\n")
	line_2 := strings.Split(body_2, "\n")
	for i := range line_1 {
		if line_1[i] == line_2[i] {
			sig = append(sig, line_1[i])
		}
	}
	return sig
}
