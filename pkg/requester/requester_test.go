package requester

import (
	"net/http"
	"testing"

	"github.com/heckintosh/fparam/pkg/utils"
)

func TestRequester(t *testing.T) {
	sampleRequest := utils.RequestPrep{
		Url:    "http://210.245.86.148/csocssrf/Level1/",
		Method: "GET",
		Header: http.Header{
			"User-Agent": {"test"},
		},
	}
	payloads := map[string]string{
		"url": "test",
	}
	resp, _ := Requester(sampleRequest, payloads)
	body_str := utils.GetRespBodyStr(resp)
	want := "<html>\n<body>\n<center><h1>THE IMAGE DOWNLOADER</h1></center>\n<br><br><br><br><br>\n<center>\n<form action='.' method='GET'>\nENTER THE IMG URL:<input type='text' name='url'>\n<br><br>\n<input type='submit'>\n</center>\n</form>\n\n</body>\n</html>"
	if body_str != want {
		t.Errorf("got %q, want %q", body_str, want)
	}
}
