package requester

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/heckintosh/fparam/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func Requester(_request utils.RequestPrep, payloads map[string]string) (*http.Response, error) {

	/*  Proxy for testing purpose
	---------------------------------------------------
	proxyurl, _ := url.Parse("http://127.0.0.1:8080")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyurl),
	}
	---------------------------------------------------
	*/

	client := http.Client{
		Timeout: 10 * time.Second,
		//Transport: transport,  (Proxy for testing purpose)
	}
	_url := _request.Url // _request stores url, method and headers that would be utilized by http package
	var req *http.Request
	var err error
	var resp *http.Response
	if _request.Method == "GET" {
		req, err = http.NewRequest("GET", _url, nil)
		q := req.URL.Query()

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Requester failed when building request.")
			return resp, err
		} else {
			req.Header = _request.Header
			for key, payload := range payloads {
				q.Add(key, payload)
			}
		}

		resp, err = client.Do(req)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Requester failed when making request.")
			return resp, err
		}

	} else if _request.Method == "JSON" {
		_request.Header["Content-Type"] = []string{"application/json"}
		jsonData, err := json.Marshal(payloads)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Requester failed when marshalling json.")
			return resp, err
		} else {
			req, err := http.NewRequest("POST", _url, bytes.NewBuffer(jsonData))

			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("Requester failed when building JSON request.")
				return resp, err
			} else {
				req.Header = _request.Header
			}

			resp, err = client.Do(req)

			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("Request failed when making JSON request")
				return resp, err
			}
		}
	} else {
		data := url.Values{}
		for key, payload := range payloads { // Turn payloads into post body
			data.Add(key, payload)
		}
		encodedData := data.Encode()
		req, err := http.NewRequest("POST", _url, strings.NewReader(encodedData))

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Requester failed when building POST request.")
			return resp, err
		} else {
			_request.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
			req.Header = _request.Header
		}

		resp, err = client.Do(req)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Request failed when making POST request")
			return resp, err
		}
	}
	return resp, nil
}
