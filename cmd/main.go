package main

import (
	"github.com/heckintosh/fparam/pkg/utils"
	"github.com/heckintosh/fparam/pkg/requester"
	"fmt"
	"strings"
	log "github.com/sirupsen/logrus"
)

func main(){
	var urls []string
	var params []string
	var method string
	requests := utils.Prepare_request()
	final_result := make(map[string]map[string]interface{})
	//for _, url := range urls{
	//	final_result[url]["params"] = params
	//	final_result[url]["method"] = method
	//	final_result[url]["headers"] =
	//}
	

	//urls := []string{"http://example.com"}
	//reqs := utils.Prepare_request(urls)
	fmt.Println(final_result, urls)
}

func initialize(_request utils.RequestPrep, wordlist []string) string{
	url := _request.Url
	if !strings.HasPrefix(url, "http"){
		log.Error("%s is not a valid url", url)
		return "skipped"
	}
	_request.Url = utils.Stable_request(url, _request.Header)
	if _request.Url != "" {
		return "skipped"
	} else {
		fuzz := utils.Random_str(6)
		fuzz_rev := utils.Reverse(fuzz)
		fuzz_map := map[string]string{fuzz : fuzz_rev}
		resp1, err1 := requester.Requester(_request, fuzz_map) //First try

		fuzz = utils.Random_str(6)
		fuzz_rev = utils.Reverse(fuzz)
		fuzz_map = map[string]string{fuzz : fuzz_rev}
		resp2, err2 := requester.Requester(_request, fuzz_map) 			// Second try

		if err1 != nil || err2 != nil{
			return "skipped"
		}
		factors := 
	}
}
