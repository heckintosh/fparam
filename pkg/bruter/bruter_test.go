package bruter

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/heckintosh/fparam/pkg/anomaly"
	"github.com/heckintosh/fparam/pkg/utils"
	"net/http"
	"sort"
	"testing"
)

func TestBruter(t *testing.T) {
	req := utils.RequestPrep{
		Url:    "http://210.245.86.148/csocssrf/Level1/",
		Method: "GET",
		Header: http.Header{
			"User-Agent": {"test"},
		},
	}
	factors := anomaly.Factors{
		SameCode:      200,
		SameBody:      "<html>\n<body>\n<center><h1>THE IMAGE DOWNLOADER</h1></center>\n<br><br><br><br><br>\n<center>\n<form action='.' method='GET'>\nENTER THE IMG URL:<input type='text' name='url'>\n<br><br>\n<input type='submit'>\n</center>\n</form>\n\n</body>\n</html>",
		SamePlaintext: "",
		LinesNum:      0,
		LinesDiff:     []string{},
		SameHeaders:   []string{"Content-Type", "Date", "Server", "Vary"},
		SameRedirect:  "/csocssrf/Level1",
		ParamMissing:  []string{"inp", "or", "in", "htm", "ty", "name", "ent", "a", "nte"},
		ValueMissing:  true,
	}
	sort.Strings(factors.ParamMissing)
	sort.Strings(factors.SameHeaders)
	param := map[string]string{"url": "111110", "htm": "111111", "in": "111112", "inp": "111113", "name": "111114"}
	got := Bruter(req, factors, param)
	spew.Dump(got)
}
