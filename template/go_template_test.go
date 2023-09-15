package template

import (
	"testing"
)

func TestGoTemplate_Create(t *testing.T) {
	type fields struct {
		Options Options
	}
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "TestGoTemplate_Create_1",
			fields: fields{Options: Options{filePath: "", handlers: []Handler{&TplHandler{}, &RunHandler{}}}},
			args: args{params: map[string]interface{}{
				"goEnv":    "https://proxy.golang.com.cn,direct",
				"appName":  "template",
				"fileName": "template.go",
				"runs":     []string{"sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories", "apk update --no-cache && apk add --no-cache tzdata"},
			}},
			wantErr: false,
		},
		{
			name:   "TestGoTemplate_Create_2",
			fields: fields{Options: Options{filePath: "", handlers: []Handler{&TplHandler{}, &RunHandler{}}}},
			args: args{params: map[string]interface{}{
				"goEnv":    "https://proxy.golang.com.cn,direct",
				"appName":  "template",
				"fileName": "template.go",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gt := &GoTemplate{
				Options: tt.fields.Options,
			}
			if err := gt.Create(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
