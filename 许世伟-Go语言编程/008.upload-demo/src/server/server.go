package main

import "tpl"
import "net/http"
import "fmt"
import "os"
import "io/ioutil"

func main() {
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/static/", handlerStatic)

	http.ListenAndServe(":8080", nil)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	tpl := tpl.NewTpl("./tpls", "tpl")
	tplDatas := make(map[string]string)

	r.ParseMultipartForm(16 * 1024 * 1024)

	tplDatas["file_list"] = ""
	tplDatas["upload_success"] = ""

	if r.MultipartForm != nil {
		if upload_file, ok := r.MultipartForm.File["upload_file"]; ok {
			for _, h := range upload_file {
				uploadFp, _ := h.Open()

				content, _ := ioutil.ReadAll(uploadFp)
				ioutil.WriteFile("./data/"+h.Filename, content, 0755)
				tmpTplDatas := make(map[string]string)
				tmpTplDatas["fname"] = h.Filename
				tplDatas["upload_success"] = tpl.Render("upload_success", tmpTplDatas)
			}
		}
	}

	dp, _ := os.Open("./data")
	dirNames, _ := dp.Readdirnames(0)
	if len(dirNames) > 0 {
		for _, dirName := range dirNames {
			if dirName[0] != '.' {
				tplDatas["file_list"] += dirName + "<br />"
			}
		}
	}

	fmt.Fprint(w, tpl.Render("index", tplDatas))
}

func handlerStatic(w http.ResponseWriter, r *http.Request) {
	fName := "./static/" + r.RequestURI[8:]
	fp, _ := os.Open(fName)

	content, _ := ioutil.ReadAll(fp)
	fmt.Fprint(w, string(content))
}
