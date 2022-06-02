package main

import (
	// "fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	// "github.com/heckintosh/fparam/pkg/anomaly"
	"github.com/heckintosh/fparam/pkg/plugins"
	"github.com/heckintosh/fparam/pkg/requester"
	"github.com/heckintosh/fparam/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func RunV1(urls []string, filepath string) map[string][]string {
	var params []string
	var url string
	wordlist := utils.GetWordList(filepath)
	requests := utils.PrepareRequest(urls)
	final_result := make(map[string][]string)
	for _, request := range requests {
		url = request.Url
		params = Initialize(request, wordlist)
		if len(params) == 0 {
			log.WithFields(log.Fields{
				"url": url,
			}).Error("Either Run failed or no params were found")
		} else {
			final_result[url] = params
		}
	}
	return final_result
}

func Initialize(_request utils.RequestPrep, wordlist []string) []string {
	url := _request.Url
	if !strings.HasPrefix(url, "http") {
		log.WithFields(log.Fields{
			"url": url,
		}).Error("Not a valid url.")
		return []string{}
	}
	_request.Url = utils.Stable_request(url, _request.Header)
	if _request.Url == "" {
		log.WithFields(log.Fields{
			"url": _request.Url,
		}).Error("Got an empty url.")
		return []string{}
	} else {
		fuzz := utils.Random_str(6)
		fuzz_rev := utils.Reverse(fuzz)
		fuzz_map := map[string]string{fuzz: fuzz_rev}
		resp1, err1 := requester.Requester(_request, fuzz_map) //First try

		fuzz = utils.Random_str(6)
		fuzz_rev = utils.Reverse(fuzz)
		fuzz_map = map[string]string{fuzz: fuzz_rev}
		_, err2 := requester.Requester(_request, fuzz_map) // Second try

		if err1 != nil || err2 != nil {
			log.WithFields(log.Fields{
				"error1": err1,
				"error2": err2,
			}).Error("Initializer failed.")
			return []string{}
		}
		body_str := utils.GetRespBodyStr(resp1)
		found := plugins.Heuristic(body_str, wordlist)
		return found
	}
}
