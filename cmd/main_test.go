package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/heckintosh/fparam/pkg/utils"
)

func TestInitalize(t *testing.T) {
	filepath, _ := filepath.Abs("../internal/db/test.txt")
	wordlist := utils.GetWordList(filepath)
	sampleRequest := utils.RequestPrep{
		Url:    "http://210.245.86.148/csocssrf/Level1/",
		Method: "GET",
		Header: http.Header{
			"User-Agent": {"test"},
		},
	}
	chunk := 300
	if len(wordlist) < chunk {
		chunk = int(len(wordlist) / 2)
	}
	test_str, test_int := Initialize(sampleRequest, wordlist, chunk)
	fmt.Println(test_str)
	fmt.Println(test_int)
}
