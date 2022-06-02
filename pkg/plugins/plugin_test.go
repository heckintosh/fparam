package plugins

import (
	"net/http"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/heckintosh/fparam/pkg/requester"
	"github.com/heckintosh/fparam/pkg/utils"
)

func TestHeuristics(t *testing.T) {
	filepath, _ := filepath.Abs("../../internal/db/test.txt")
	wordlist := utils.GetWordList(filepath)
	reqprep := utils.RequestPrep{
		Url:    "http://210.245.86.148/csocssrf/Level1",
		Method: "GET",
		Header: http.Header{
			"User-Agent": {"test"},
		},
	}
	payloads := map[string]string{
		"test": "test",
	}
	resp, _ := requester.Requester(reqprep, payloads)
	resp_body_str := utils.GetRespBodyStr(resp)
	spew.Config.Dump(resp_body_str)
	got := Heuristic(resp_body_str, wordlist)
	want := []string{"url"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
