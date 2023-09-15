package config

import (
	"reflect"
	"testing"
)

func TestNewGod(t *testing.T) {
	tests := []struct {
		name string
		want *God
	}{
		{
			name: "TestNewGod_1",
			want: &God{
				WorkDir:  "C:\\Users\\leig\\Developer\\github\\god\\config",
				File:     "config.go",
				FileType: GoType,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGod(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_map2String(t *testing.T) {
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name      string
		args      args
		wantParam string
	}{
		{
			name:      "Test_map2String_1",
			args:      args{map[string]string{"a": "b", "c": "d"}},
			wantParam: "a=b,c=d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParam := map2String(tt.args.params); gotParam != tt.wantParam {
				t.Errorf("map2String() = %v, want %v", gotParam, tt.wantParam)
			}
		})
	}
}
