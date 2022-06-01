package main

import (
	"fmt"
	"strings"

	"github.com/heckintosh/fparam/pkg/anomaly"
	"github.com/heckintosh/fparam/pkg/plugins"
	"github.com/heckintosh/fparam/pkg/requester"
	"github.com/heckintosh/fparam/pkg/utils"
	log "github.com/sirupsen/logrus"
)

//func main() {
//	var urls []string
//	var params []string
//	var method string
//	requests := utils.Prepare_request()
//	final_result := make(map[string]map[string]interface{})
//	//for _, url := range urls{
//	//	final_result[url]["params"] = params
//	//	final_result[url]["method"] = method
//	//	final_result[url]["headers"] =
//	//}
//
//	//urls := []string{"http://example.com"}
//	//reqs := utils.Prepare_request(urls)
//	fmt.Println(final_result, urls)
//}

//type Dictionary map[string]interface{}
//
//func narrower(_request utils.RequestPrep, factors anomaly.Factors, param_groups []Dictionary){
//	anomalous_params := []string{}
//	for _, param range := range param_groups{
//		go bruter.Bruter(_request, factors, param_groups)
//	}
//}

func Initialize(_request utils.RequestPrep, wordlist []string, chunks int) (string, int) {
	url := _request.Url
	if !strings.HasPrefix(url, "http") {
		log.WithFields(log.Fields{
			"url": url,
		}).Error("Not a valid url.")
		return "skipped", 0
	}
	_request.Url = utils.Stable_request(url, _request.Header)
	if _request.Url != "" {
		return "skipped", 0
	} else {
		fuzz := utils.Random_str(6)
		fuzz_rev := utils.Reverse(fuzz)
		fuzz_map := map[string]string{fuzz: fuzz_rev}
		resp1, err1 := requester.Requester(_request, fuzz_map) //First try

		fuzz = utils.Random_str(6)
		fuzz_rev = utils.Reverse(fuzz)
		fuzz_map = map[string]string{fuzz: fuzz_rev}
		resp2, err2 := requester.Requester(_request, fuzz_map) // Second try

		if err1 != nil || err2 != nil {
			return "skipped", 0
		}
		factors := anomaly.Define(resp1, resp2, fuzz, fuzz_rev, wordlist)
		found := plugins.Heuristic(utils.GetRespBodyStr(resp1), wordlist)

		populated := utils.Populate(wordlist)
		param_groups := utils.Slicer(populated, int(len(wordlist)/chunks))
		fmt.Println(factors)
		fmt.Println(found)
		return "", len(param_groups)
	}
}
