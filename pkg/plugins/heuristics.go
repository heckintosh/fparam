package plugins

import (
	"regexp"

	"github.com/davecgh/go-spew/spew"
	"github.com/heckintosh/fparam/pkg/utils"
)

func is_not_junk(param string) bool {
	r, _ := regexp.Compile("^[A-Za-z0-9_]+$")
	isMatched := r.MatchString(param)
	return isMatched
}

func Heuristic(resp_body string, wordlist []string) []string {
	potential_params := []string{}

	str1 := `(?i)<input.+?name=["']?([^"'\s>]+)`
	r_input_names := regexp.MustCompile(str1)

	str2 := `(?i)<input.+?id=["']?([^"'\s>]+)`
	r_input_ids := regexp.MustCompile(str2)

	str3 := `(?:[;\n]|\bvar|\blet)(\w+)\s*=\s*(?:['"]{1,2}|true|false|null)`
	r_empty_vars := regexp.MustCompile(str3)

	str4 := `['"](\w+?)['"]\s*:\s*['"]`
	r_map_keys := regexp.MustCompile(str4)

	input_names := r_input_names.FindStringSubmatch(resp_body)
	potential_params = append(potential_params, input_names...)

	input_ids := r_input_ids.FindStringSubmatch(resp_body)
	potential_params = append(potential_params, input_ids...)

	var empty_vars []string
	var map_keys []string
	for _, script := range utils.ExtractJs(resp_body) {
		empty_vars = r_empty_vars.FindStringSubmatch(script)
		potential_params = append(potential_params, empty_vars...)

		map_keys = r_map_keys.FindStringSubmatch(script)
		potential_params = append(potential_params, map_keys...)
	}

	if len(potential_params) == 0 {
		return []string{}
	}

	found := []string{}
	spew.Dump(found)
	for _, word := range potential_params {
		if is_not_junk(word) && !utils.Contains(word, found) {
			found = append(found, word)
			if utils.Contains(word, wordlist) {
				for i, v := range wordlist {
					if v == word {
						wordlist = append(wordlist[:i], wordlist[i+1:]...)
						break
					}
				}
			}
			wordlist = append([]string{word}, wordlist...)
		}
	}
	return found
}
