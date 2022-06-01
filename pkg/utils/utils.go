package utils

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
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

func Extract_headers(headers string) map[string]string {
	tmp := strings.ReplaceAll(headers, "\\n", "\n")
	return Parse_headers(tmp)
}

func Parse_headers(header string) map[string]string {
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

func Populate(array []string) map[string]string {
	//converts a list of params into param and value pair
	result := make(map[string]string)
	for i := range array {
		result[array[i]] = strings.Repeat("1", 6-len(strconv.Itoa(i))) + strconv.Itoa(i)
	}
	return result
}


func Diff_map(body_1 string, body_2 string) []string {
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

func ExtractJs(resp_str string) []string {
	r1 := regexp.MustCompile(regexp.QuoteMeta(`(?i)<script[> ]`))
	r2 := regexp.MustCompile(regexp.QuoteMeta(`(?i)</script>`))
	split := r1.Split(resp_str, -1)
	scripts := []string{}
	var actual_parts []string
	for _, part := range split {
		actual_parts = r2.Split(part, 2)
		if len(actual_parts) > 1 {
			scripts = append(scripts, actual_parts[0])
		}
	}
	return scripts
}

func Contains(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}
	return false
}

func GetRespBodyStr(resp *http.Response) string {
	body := resp.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}

func GetWordList(filepath string) []string {
	file, err := os.Open(filepath)
	var words []string
	if err != nil {
		log.Fatal(err)
	}
	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		words = append(words, Scanner.Text())
	}
	return words
}

func Confirm(array_dict []Dictionary, usable []Dictionary) []Dictionary{
	params_groups := [] Dictionary{}
	for i := range array_dict{
		if(len(array_dict[i]) == 1){
			usable = append(usable, array_dict[i])
		} else{
			params_groups = append(params_groups,array_dict[i])
		}
	}
	return params_groups
}






type Array_Dict struct{
	Dicts []Dictionary
}


type Dictionary map[string]interface{}
