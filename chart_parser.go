package gchart

import (
	"errors"
	"github.com/zieckey/goini"
	"os"
	"path/filepath"
)

const DataPrefix = "Data|"

type ChartIf interface {
	Parse(ini *goini.INI, date string) (map[string]string, error)
	Template() string
}

type TemplateArgs struct {
	args map[string]string
	tmpl string
}

func Parse(file string, date string) (tt TemplateArgs, err error) {
	ini := goini.New()
	err = ini.ParseFile(file)
	if err != nil {
		return tt, err
	}

	if f, ok := ChartHandlers["column"]; ok {
		tt.args, err = f.Parse(ini, date)
		tt.tmpl = f.Template()
	}

	return tt, err
}

func LookupChartFiles(dir string) ([]string, error) {
	var files []string = make([]string, 0, 5)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if ok, err := filepath.Match("*.chart", f.Name()); err != nil {
			return err
		} else if ok {
			files = append(files, path)
		}
		return nil
	})

	if len(files) == 0 {
		return files, errors.New("Not found any *.chart files")
	}

	return files, err
}
