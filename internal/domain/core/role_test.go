package core

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/rendau/dutchman/internal/domain/entities"
)

func TestRole_parseRemoteJson(t *testing.T) {
	cr := &Role{}

	tests := []struct {
		src  []byte
		path string
		want []*entities.RoleRemoteRepItemSt
	}{
		{
			src: []byte(`{
				"perms": [
					{"code": "c1", "is_all": true, "dsc": "c1 desc"},
					{"code": "c2", "is_all": false, "dsc": "c2 desc"}
				]
			}`),
			path: "   perms   ",
			want: []*entities.RoleRemoteRepItemSt{
				{Code: "c1", Dsc: "c1 desc"},
				{Code: "c2", Dsc: "c2 desc"},
			},
		},
		{
			src: []byte(`[
				{"code": "c1", "is_all": true, "dsc": "c1 desc"},
				{"code": "c2", "is_all": false, "dsc": "c2 desc"}
			]`),
			path: "",
			want: []*entities.RoleRemoteRepItemSt{
				{Code: "c1", Dsc: "c1 desc"},
				{Code: "c2", Dsc: "c2 desc"},
			},
		},
		{
			src: []byte(`{
				"k1": {
					"foo": "bar",
					"k2": {
						"k3": {
							"perms": [
								{"code": "c1", "is_all": true, "dsc": "c1 desc"},
								{"code": "c2", "is_all": false, "dsc": "c2 desc"}
							],
							"foo": "bar"
						}
					}
				}
			}`),
			path: "..k1.k2 .  k3.perms .",
			want: []*entities.RoleRemoteRepItemSt{
				{Code: "c1", Dsc: "c1 desc"},
				{Code: "c2", Dsc: "c2 desc"},
			},
		},
	}
	for ttI, tt := range tests {
		t.Run(strconv.Itoa(ttI+1), func(t *testing.T) {
			got := cr.parseRemoteJson(tt.src, tt.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRemoteJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
