package template

import (
	_ "embed"
	"strings"
)

//go:embed go_dockerfile.tpl
var defaultGoTemplate string

type Template interface {
	Create(params map[string]interface{}) (err error)
}

type Options struct {
	filePath string
	handlers []Handler
}

type Option func(options *Options)

func WithFilePath(filePath string) Option {
	return func(options *Options) {
		if strings.EqualFold(filePath, "") {
			return
		}
		options.filePath = filePath
	}
}

func WithHandlers(handlers ...Handler) Option {
	return func(options *Options) {
		if len(handlers) > 0 {
			options.handlers = handlers
		}
	}
}
