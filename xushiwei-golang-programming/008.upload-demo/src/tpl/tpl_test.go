package tpl

import "testing"
import "fmt"

func TestNewTpl(t *testing.T) {
	tpl := NewTpl("./tpls", "tpl")
	datas := make(map[string]string)
	datas["uname"] = "goosman-lei"
	datas["uid"] = "5012470"

	result := tpl.Render("info", datas)
	fmt.Println(result)
}
