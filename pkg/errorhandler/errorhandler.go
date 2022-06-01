package errorhandler

import (
	"errors"
	"github.com/heckintosh/fparam/pkg/anomaly"
	"net/http"
)

func errorhandler(resp *http.Response, err error, factors anomaly.Factors) {
	status_slice := []int{400, 413, 418, 429, 503}
	_, found := Find(status_slice, resp.StatusCode)
	if found {
		if Find([]int{429, 418}, resp.StatusCode) {

		}
	}

}

func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
