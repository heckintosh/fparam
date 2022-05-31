package requester

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/heckintosh/fparam/pkg/utils"
)

func TestRequester(t *testing.T) {
	sampleRequest := utils.RequestPrep{
		Url:    "http://example.com",
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

	count := strings.Count(body_str, "\n")
	fmt.Printf("Count: %v\n", count)
}
