package core

import (
	"reflect"
	"testing"

	"github.com/rendau/dutchman/internal/domain/entities"
)

func TestRole_parseRemoteJson(t *testing.T) {
	type fields struct {
		r *St
	}
	type args struct {
		src  []byte
		path []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.RoleRemoteRepItemSt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Role{
				r: tt.fields.r,
			}
			got, err := c.parseRemoteJson(tt.args.src, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRemoteJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRemoteJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
