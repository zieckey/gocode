package strings

import (
	"testing"
	"strings"
)

func filter(r rune) bool {
	filters := "abc123，。《》“”"
	if strings.IndexRune(filters, r) >= 0 {
		return true
	}

	return false
}

func TestRemoveIf(t * testing.T) {
	testCases := []struct {
		r, want string
	} {
		{"《白鹿原》","白鹿原"},
		{"《白鹿原》以陕西关中平原上素有“仁义村”之称的白鹿村为背景，展现了白姓和鹿姓两大家族祖孙三代的恩怨纷争。","白鹿原以陕西关中平原上素有仁义村之称的白鹿村为背景展现了白姓和鹿姓两大家族祖孙三代的恩怨纷争"},
		{"hello world","hello world"},
		{"1234567890","4567890"},
		{"remove all","remove ll"},
	}

	for i, c := range testCases {
		e := RemoveIf(c.r, filter)
		if e != c.want {
			t.Errorf("[%v] test failed", testCases[i].r)
		}
	}

}


