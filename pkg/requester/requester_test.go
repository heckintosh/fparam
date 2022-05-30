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
	got, _ := Requester(sampleRequest, payloads)
	got_body := got.Body

	buf := new(bytes.Buffer)
	buf.ReadFrom(got_body)

	got_str := buf.String()
	count := strings.Count(got_str, "\n")
	fmt.Printf("Count: %v\n", count)
}
