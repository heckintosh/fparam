package plugins

import (
	"net/http"
	"regexp"
)

func is_not_junk(param string) bool {
	r, _ := regexp.Compile("^[A-Za-z0-9_]+$")
	isMatched := r.MatchString(param)
	return isMatched
}

func heuristic(resp *http.Response, wordlist []string) {
	potential_params := []string{}

	r_input_names, _ := regexp.Compile(regexp.QuoteMeta(`(?i)<input.+?name=["']?([^"'\s>]+)`))
	r_input_ids, _ := regexp.Compile(regexp.QuoteMeta(`(?i)<input.+?id=["']?([^"'\s>]+))`))
	r_empty_vars, _ := regexp.Compile(regexp.QuoteMeta(`(?:[;\n]|\bvar|\blet)(\w+)\s*=\s*(?:['"]{1,2}|true|false|null)`))
	r_map_keys, _ = regexp.Compile(regexp.QuoteMeta(`['"](\w+?)['"]\s*:\s*['"]`))

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body_str := buf.String()

	input_names := r_input_names.FindAllString(body_str, -1)
	potential_params = append(potential_params, input_names...)

	input_ids := r_input_ids.FindAllString(body_str, -1)
	potential_params = append(potential_params, input_ids...)

}
