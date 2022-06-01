package anomaly

import (
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/heckintosh/fparam/pkg/utils"
)

type Factors struct {
	SameCode      int
	SameBody      string
	SamePlaintext string
	LinesNum      int
	LinesDiff     []string
	SameHeaders   []string
	SameRedirect  string
	ParamMissing  []string
	ValueMissing  bool
}

func New() Factors {
	return Factors{
		SameCode:      0,
		SameBody:      "",
		SamePlaintext: "",
		LinesNum:      0,
		LinesDiff:     []string{},
		SameHeaders:   []string{},
		SameRedirect:  "",
		ParamMissing:  []string{},
		ValueMissing:  false,
	}
}

func Define(resp1 *http.Response, resp2 *http.Response, param string, value string, wordlist []string) Factors {
	f := New()
	body1 := utils.GetRespBodyStr(resp1)
	body2 := utils.GetRespBodyStr(resp2)

	if resp1.StatusCode == resp2.StatusCode {
		f.SameCode = resp1.StatusCode
	}

	key_resp1_header := utils.GetKeys(resp1.Header)
	key_resp2_header := utils.GetKeys(resp2.Header)
	if utils.IsEqual(key_resp1_header, key_resp2_header) {
		header_keys := []string{}
		for key := range resp1.Header {
			header_keys = append(header_keys, key)
		}
		sort.Strings(header_keys)
		f.SameHeaders = header_keys
	}

	parse1, _ := url.Parse(resp1.Request.URL.String())
	parse2, _ := url.Parse(resp2.Request.URL.String())
	if parse1.Path == parse2.Path {
		f.SameRedirect = parse1.Path
	} else {
		f.SameRedirect = ""
	}

	if reflect.DeepEqual(body1, body2) {

		f.SameBody = body1
	} else if strings.Count(body1, "\n") == strings.Count(body2, "\n") {
		f.LinesNum = strings.Count(body1, "\n")
	} else if utils.Remove_tags(body1) == utils.Remove_tags(body2) {
		f.SamePlaintext = utils.Remove_tags(body1)
	}

	if !strings.Contains(body2, param) {
		for _, word := range wordlist {
			if strings.Contains(body2, word) {
				f.ParamMissing = append(f.ParamMissing, word)
			}
		}
		sort.Strings(f.ParamMissing)
	}
	if !strings.Contains(body2, value) {
		f.ValueMissing = true
	}
	return f
}

func Compare(resp *http.Response, factors Factors, params map[string]string) (string, map[string]string) {
	these_header := utils.GetKeys(resp.Header)
	sort.Strings(these_header)
	resp_body_str := utils.GetRespBodyStr(resp)
	sort.Strings(these_header)
	if factors.SameCode != 0 && resp.StatusCode != factors.SameCode {
		return "http code", params
	}
	if len(factors.SameHeaders) != 0 && reflect.DeepEqual(these_header, factors.SameHeaders) {
		return "http headers", params
	}
	if len(factors.SameBody) != 0 && resp_body_str != factors.SameBody {
		return "body length", params
	}
	if factors.LinesNum != 0 && strings.Count(resp_body_str, "\n") != factors.LinesNum {
		return "number of lines", params
	}
	if len(factors.SamePlaintext) != 0 && utils.Remove_tags(resp_body_str) != factors.SamePlaintext {
		return "text length", params
	}
	if len(factors.ParamMissing) != 0 {
		for param, _ := range params {
			if len(param) < 5 {
				continue
			}
			reg_str := `[\'"\s]` + param + `[\'"\s]`
			r := regexp.MustCompile(reg_str)
			matched := r.MatchString(resp_body_str)
			if !utils.Contains(param, factors.ParamMissing) && matched {
				return "param name reflection", params
			}
		}
	}
	if factors.ValueMissing {
		for _, value := range params {
			reg_str := `[\'"\s]` + value + `[\'"\s]`
			r := regexp.MustCompile(reg_str)
			matched := r.MatchString(resp_body_str)
			if strings.Contains(value, resp_body_str) && matched {
				return "param value reflection", params
			}
		}
	}
	return "", map[string]string{}
}
