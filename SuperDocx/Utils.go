package docx

import (
	"strings"
)

func index(Base string, search string) int {
	if Base == search {
		return 0
	}
	baseSplit := strings.Split(Base, "")
	searchSplit := strings.Split(search, "")
	if len(baseSplit) < len(searchSplit) {
		return -1
	}
	if len(searchSplit) == 0 {
		return -1
	}
	if len(baseSplit) == 0 {
		return -1
	}
	finded := -1
	for baseIndex, baseChar := range baseSplit {
		if baseChar != searchSplit[0] {
			continue
		}
		find := true
		for searchIndex, searchChar := range searchSplit {
			if baseIndex+searchIndex >= len(baseSplit) {
				return -1
			}
			if baseSplit[baseIndex+searchIndex] != searchChar {
				find = false
				break
			}
		}
		if find == true {
			finded = baseIndex
			break
		}
	}
	return finded

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
