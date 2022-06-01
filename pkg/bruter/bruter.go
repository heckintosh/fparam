package bruter

import (
	"github.com/heckintosh/fparam/pkg/anomaly"
	"github.com/heckintosh/fparam/pkg/requester"
	"github.com/heckintosh/fparam/pkg/utils"
)

func Bruter(_request utils.RequestPrep, factors anomaly.Factors, params map[string]interface{}) map[string]string {
	resp, _ := requester.Requester(_request, params)
	_, comparison_result := anomaly.Compare(resp, factors, params)
	return comparison_result
}
