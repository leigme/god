package template

import (
	"bytes"
	"fmt"
	"github.com/leigme/god/config"
	"os"
	"text/template"
)

type Handler interface {
	Handle(oldContent string, params map[string]interface{}) (newContent string)
}
type TplHandler struct {
}

func (th *TplHandler) Handle(oldContent string, params map[string]interface{}) (newContent string) {
	tmpl, err := template.New(config.Dockerfile).Parse(oldContent)
	if err != nil {
		return oldContent
	}
	buffer := bytes.NewBufferString(newContent)
	err = tmpl.Execute(buffer, params)
	if err != nil {
		return oldContent
	}
	return buffer.String()
}

type RunHandler struct {
}

func (rh *RunHandler) Handle(oldContent string, params map[string]interface{}) (newContent string) {
	if runsArr, ok := params["runs"].([]string); ok {
		runs := runsArr2Map(runsArr)
		return os.Expand(oldContent, func(key string) string {
			return runs[key]
		})
	}
	return oldContent
}

func runsArr2Map(runs []string) map[string]string {
	m := make(map[string]string)
	for k, v := range runs {
		m[fmt.Sprintf("run%d", k)] = v
	}
	return m
}
