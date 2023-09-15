package template

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/leigme/god/config"
	"github.com/leigme/god/utils"
	"io"
	"log"
	"os"
	"strings"
)

type GoTemplate struct {
	Options
}

func (gt *GoTemplate) Create(params map[string]interface{}) (err error) {
	if utils.CheckFileExist(config.Dockerfile) {
		return errors.New(fmt.Sprintf("file %s already exists", config.Dockerfile))
	}
	tpl := defaultGoTemplate
	if !strings.EqualFold(gt.filePath, "") {
		r, err := os.Open(gt.filePath)
		if err != nil {
			return err
		}
		all, err := io.ReadAll(r)
		if err != nil {
			return err
		}
		tpl = string(all)
	}
	for _, handler := range gt.handlers {
		tpl = handler.Handle(tpl, params)
	}
	err = createFile(tpl)
	if err != nil {
		e := deleteFile()
		log.Println(e)
	}
	return err
}

func NewGoTemplate(opts ...Option) Template {
	options := Options{}
	for _, apply := range opts {
		apply(&options)
	}
	gt := &GoTemplate{
		options,
	}
	return gt
}

func createFile(tpl string) error {
	dockerfile, err := os.Create(config.Dockerfile)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(dockerfile)
	bw := bufio.NewWriter(dockerfile)
	_, err = bw.WriteString(tpl)
	if err != nil {
		return err
	}
	err = bw.Flush()
	if err != nil {
		return err
	}
	return nil
}

func deleteFile() error {
	return os.Remove(config.Dockerfile)
}
