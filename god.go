/*
Copyright © 2023 NAME HERE leigme@gmail.com
*/
package main

import (
	"github.com/leigme/god/config"
	"github.com/leigme/god/template"
	"log"
)

func main() {
	god := config.NewGod()
	log.Println(god)
	createDockerfile(god)
}

func createDockerfile(god *config.God) {
	var t template.Template
	switch god.FileType() {
	case config.GoType:
		t = template.NewGoTemplate(template.WithFilePath(god.TemplatePath()), template.WithHandlers(&template.TplHandler{}, &template.RunHandler{}))
		break
	case config.JavaType:
		break
	}
	err := t.Create(god.Params())
	if err != nil {
		return
	}
}
