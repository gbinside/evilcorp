package censorship

import (
	"regexp"
	"strings"

)

func Replace(s string) string {
	words := map[string]string{
		"bad":       "ungood",
		"better":    "gooder",
		"objection": "thoughtcrime",
		"agree":     "crimestop",
	}

	var keys []string
	for k := range words {
		keys = append(keys, k)
	}

	regexpattern := strings.Join(keys, "|")

	re := regexp.MustCompile("(?i)" + regexpattern)
	ret := re.ReplaceAllStringFunc(s, func(m string) string {
		w := words[strings.ToLower(m)]
		if m[0] < 91 && m[0] > 64 {
			w = strings.ToUpper(w[:1]) + w[1:]
		}
		return w
	})

	return ret
}

func Censorship(s string) string {

	s2 := Replace(s)

	words := "nice|pony|sun|light|fun|happy|funny|joy|friend"
	words_list := strings.Split(words, "|")
	ars := make([]string, len(words_list))
	for i, w := range words_list {
		ars[i] = w + "[a-zA-Z!]*"
	}
	regexpattern := strings.Join(ars, "|")

	re := regexp.MustCompile(regexpattern)
	ret := re.ReplaceAllStringFunc(s2, func(m string) string {
		return strings.Repeat("X", len(m))
	})

	return ret
}
