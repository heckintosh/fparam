package anomaly

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/heckintosh/fparam/pkg/utils"
	"net/http"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestDefine(t *testing.T) {
	//proxyurl, _ := url.Parse("http://127.0.0.1:8080")
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxyurl),
	//}
	url := "http://210.245.86.148/csocssrf/Level1/"
	client := http.Client{
		Timeout: 10 * time.Second,
		//Transport: transport,
	}
	header := http.Header{
		"User-Agent":                {"Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0"},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
		"Accept-Language":           {"en-US,en;q=0.5"},
		"Connection":                {"close"},
		"Upgrade-Insecure-Requests": {"1"},
	}
	req1, _ := http.NewRequest("GET", url, nil)
	req2, _ := http.NewRequest("GET", url, nil)
	req1.Header = header
	req2.Header = header
	resp1, _ := client.Do(req1)
	resp2, _ := client.Do(req2)
	resp_body_str := utils.GetRespBodyStr(resp1)
	param := "027009"
	value := "900720"
	wordlist := []string{"ent", "inp", "nte", "ty", "name", "htm", "a", "or", "in"}

	got_factors := Define(resp1, resp2, param, value, wordlist)
	got_factors.SameBody = resp_body_str
	sort.Strings(got_factors.ParamMissing)

	want := Factors{
		SameCode:      200,
		SameBody:      "<html>\n<body>\n<center><h1>THE IMAGE DOWNLOADER</h1></center>\n<br><br><br><br><br>\n<center>\n<form action='.' method='GET'>\nENTER THE IMG URL:<input type='text' name='url'>\n<br><br>\n<input type='submit'>\n</center>\n</form>\n\n</body>\n</html>",
		SamePlaintext: "",
		LinesNum:      0,
		LinesDiff:     []string{},
		SameHeaders:   []string{"Content-Type", "Date", "Server", "Vary"},
		SameRedirect:  "/csocssrf/Level1/",
		ParamMissing:  []string{"in", "ty", "inp", "htm", "name", "or", "nte", "a", "ent"},
		ValueMissing:  true,
	}
	sort.Strings(want.ParamMissing)

	if !reflect.DeepEqual(got_factors, want) {
		fmt.Println("------------------------------------------------------")
		spew.Dump(got_factors)
		fmt.Println("------------------------------------------------------")
		spew.Dump(want)
		t.Error("Define test failed. The value below are different")
	}
}
