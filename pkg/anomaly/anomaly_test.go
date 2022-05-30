package anomaly

import (
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"testing"
)

func TestDefine(t *testing.T) {
	resp1, _ := http.Get("http://210.245.86.148/csocssrf/Level1")
	resp2, _ := http.Get("http://210.245.86.148/csocssrf/Level1")
	param := "027009"
	value := "900720"
	wordlist := []string{"ent", "inp", "nte", "ty", "name", "htm", "a", "or", "in"}
	test := Define(resp1, resp2, param, value, wordlist)
	spew.Dump(test)
}
