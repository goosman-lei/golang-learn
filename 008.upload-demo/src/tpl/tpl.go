package tpl

import "os"
import "io/ioutil"
import "regexp"

type Tpl struct {
	tplPath   string
	tplSuffix string
}

func NewTpl(tplPath, tplSuffix string) (tpl *Tpl) {
	_, err := os.Stat(tplPath)
	if err != nil {
		os.Mkdir(tplPath, 0755)
	}

	tpl = new(Tpl)
	tpl.tplPath = tplPath
	tpl.tplSuffix = tplSuffix

	return
}

func (tpl *Tpl) Render(tplName string, data map[string]string) (result string) {
	tplFname := tpl.tplPath + "/" + tplName + "." + tpl.tplSuffix

	tplFp, err := os.Open(tplFname)
	defer tplFp.Close()
	if err != nil {
		panic(err)
	}

	tplContent, err := ioutil.ReadAll(tplFp)
	if err != nil {
		panic(err)
	}
	tplContentStr := string(tplContent)

	r, err := regexp.Compile("\\{%[\\w-]+%\\}")
	result = r.ReplaceAllStringFunc(tplContentStr, func(info string) string {
		key := info[2 : len(info)-2]
		return data[key]
	})

	return
}
