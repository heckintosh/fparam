package anomaly

import (
	"bytes"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"github.com/heckintosh/fparam/pkg/utils"
)

type Factors struct {
	SameCode      string
	SameBody      string
	SamePlaintext string
	LinesNum      int
	SameHeaders   []string
	SameRedirect  string
	ParamMissing  []string
	ValueMissing  bool
}

func Define(resp1 *http.Response, resp2 *http.Response, param string, value string, wordlist []string) Factors {
	f := Factors{}
	body1 := resp1.Body
	body2 := resp2.Body

	buf := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)

	buf.ReadFrom(body1)
	body1_str := buf.String()
	buf2.ReadFrom(body2)
	body2_str := buf.String()

	if resp1.Status == resp2.Status {
		f.SameCode = resp1.Status
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

	if reflect.DeepEqual(body1_str, body2_str) {

		f.SameBody = body1_str
	} else if strings.Count(body1_str, "/n") == strings.Count(body2_str, "/n") {
		f.LinesNum = strings.Count(body1_str, "/n")
	} else if utils.Remove_tags(body1_str) == utils.Remove_tags(body2_str) {
		f.SamePlaintext = utils.Remove_tags(body1_str)
	}

	if !strings.Contains(body2_str, param) {
		for _, word := range wordlist {
			if strings.Contains(body2_str, word) {
				f.ParamMissing = append(f.ParamMissing, word)
			}
		}
	}
	if !strings.Contains(body2_str, value) {
		f.ValueMissing = true
	}
	return f
}
