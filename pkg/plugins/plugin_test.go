package plugins

import (
	"net/http"
	"path/filepath"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/heckintosh/fparam/pkg/requester"
	"github.com/heckintosh/fparam/pkg/utils"
)

func TestHeuristics(t *testing.T) {
	filepath, _ := filepath.Abs("../../cmd/large.txt")
	wordlist := utils.GetWordList(filepath)
	reqprep := utils.RequestPrep{
		Url:    "https://binance.com/en/trade/WAVES_USDT",
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
	got := Heuristic(resp_body_str, wordlist)
	//want := []string{"url"}
	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf("got %q, wanted %q", got, want)
	//}
	spew.Dump(got)
}
