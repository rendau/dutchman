package util

import (
	"strconv"
	"testing"
)

func TestStrTrimCommonSuffix(t *testing.T) {
	tests := []struct {
		a     string
		b     string
		want1 string
		want2 string
		want3 string
	}{
		{"", "", "", "", ""},
		{"/a/b", "/a/b", "", "", "/a/b"},
		{"", "/a/b", "", "/a/b", ""},
		{"/a/b", "", "/a/b", "", ""},
		{"/a/b/c", "/b/c", "/a", "", "/b/c"},
		{"/b/c", "/a/b/c", "", "/a", "/b/c"},
		{"/a/b/c", "/x/y/c", "/a/b", "/x/y", "/c"},
		{"/a/b/c", "/z/x/y", "/a/b/c", "/z/x/y", ""},
	}
	for ttI, tt := range tests {
		t.Run(strconv.Itoa(ttI+1), func(t *testing.T) {
			got, got1, got2 := StrTrimCommonSuffix(tt.a, tt.b)
			if got != tt.want1 {
				t.Errorf("StrTrimCommonSuffix() got1 = %v, want %v", got, tt.want1)
			}
			if got1 != tt.want2 {
				t.Errorf("StrTrimCommonSuffix() got2 = %v, want %v", got1, tt.want2)
			}
			if got2 != tt.want3 {
				t.Errorf("StrTrimCommonSuffix() got3 = %v, want %v", got2, tt.want3)
			}
		})
	}
}
