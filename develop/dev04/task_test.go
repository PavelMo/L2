package main

import (
	"strings"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	answers := map[string][]string{
		"диктор": {"диктор", "дротик", "кордит"},
		"нищета": {"нищета", "тщание", "щетина"},
		"портик": {"портик", "портки", "приток", "тропик"},
	}
	dictionary := []string{
		"Диктор",
		"дротик",
		"Дротик",
		"кордит",
		"Нищета",
		"тЩАние",
		"Щетина",
		"Портик",
		"портки",
		"приток",
		"тропик",
	}
	res := findAnagrams(dictionary)
	for key, val := range res {
		var (
			anagrams []string
			ok       bool
		)
		if anagrams, ok = answers[key]; !ok {
			t.Errorf("excess key: %s", key)
		}
		joinedResultAnagrams := strings.Join(anagrams, " ")
		joinedMyAnagrams := strings.Join(val, " ")
		if joinedMyAnagrams != joinedResultAnagrams {
			t.Errorf("wrong anagrams:\nShould: %s\nGot: %s\n", joinedResultAnagrams, joinedMyAnagrams)
		}
	}
}
