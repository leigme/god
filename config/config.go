package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Kind int

const (
	Dockerfile      = "Dockerfile"
	DefaultTemplate = "go_dockerfile.tpl"
)

const (
	GoType Kind = iota + 1
	JavaType
)

type God struct {
	workDir      string
	fileName     string
	fileType     Kind
	templatePath string
	params       string
	goMod        bool
	goSum        bool
	cmd          string
}

func (god *God) FileName() string {
	return god.fileName
}

func (god *God) AppName() string {
	return strings.TrimSuffix(god.fileName, ".go")
}

func (god *God) Params() (params map[string]interface{}) {
	params = make(map[string]interface{})
	if strings.EqualFold(god.params, "") {
		return
	}
	temps := strings.Split(god.params, ";")
	for _, temp := range temps {
		if strings.Contains(temp, "=") {
			kv := strings.SplitN(temp, "=", 2)
			params[kv[0]] = kv[1]
		}
	}
	return
}

func (god *God) FileType() Kind {
	return god.fileType
}

func (god *God) TemplatePath() string {
	return god.templatePath
}

func NewGod() *God {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	god := &God{
		workDir: dir,
	}
	flag.StringVar(&god.fileName, "o", "", "")
	flag.StringVar(&god.templatePath, "t", "", "")
	flag.StringVar(&god.params, "p", "", "")
	flag.StringVar(&god.cmd, "c", "", "")
	flag.Parse()
	defaultFile(god)
	defaultTemplate(god)
	defaultParams(god)
	return god
}

func readDirFiles() (files []string) {
	files = make([]string, 0)
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	dirFiles, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	if len(dirFiles) == 0 {
		return
	}
	for _, df := range dirFiles {
		files = append(files, df.Name())
	}
	return
}

func defaultFile(god *God) {
	files := readDirFiles()
	if strings.EqualFold(god.fileName, "") {
		if len(files) == 0 {
			log.Fatal("no files")
		}
		goMap := make([]string, 0)
		javaMap := make([]string, 0)
		for _, f := range files {
			if strings.HasSuffix(f, "_test.go") {
				break
			}
			if strings.EqualFold(f, "go.mod") {
				god.goMod = true
			}
			if strings.EqualFold(f, "go.sum") {
				god.goSum = true
			}
			if strings.HasSuffix(f, ".go") {
				goMap = append(goMap, filepath.Join(god.fileName, f))
			}
			if strings.HasSuffix(f, ".java") {
				javaMap = append(javaMap, filepath.Join(god.fileName, f))
			}
		}
		if len(goMap) == 1 {
			god.fileName = goMap[0]
		} else if len(javaMap) == 1 {
			god.fileName = javaMap[0]
		} else {
			log.Fatal("no find only golang file or java file")
		}
	}
	if strings.HasSuffix(god.fileName, ".go") {
		god.fileType = GoType
	} else if strings.HasSuffix(god.fileName, ".java") {
		god.fileType = JavaType
	}
}

func defaultTemplate(god *God) {
	if strings.EqualFold(god.templatePath, "") {
		files := readDirFiles()
		if len(files) < 1 {
			log.Fatal("no template found")
		}
		for _, f := range files {
			if strings.EqualFold(f, DefaultTemplate) {
				god.templatePath = f
			}
		}
	}
}

func defaultParams(god *God) {
	if strings.EqualFold(god.params, "") {
		dp := make(map[string]string)
		dp["goEnv"] = os.Getenv("GOPROXY")
		dp["appName"] = god.AppName()
		dp["fileName"] = god.FileName()
		if god.goMod {
			dp["goMod"] = "true"
		}
		if god.goSum {
			dp["goSum"] = "true"
		}
		god.params = map2String(dp)
	}
}

func map2String(params map[string]string) (param string) {
	for k, v := range params {
		param = fmt.Sprintf("%s;%s=%s", param, k, v)
	}
	param = strings.TrimPrefix(param, ";")
	param = strings.TrimSuffix(param, ";")
	return
}
