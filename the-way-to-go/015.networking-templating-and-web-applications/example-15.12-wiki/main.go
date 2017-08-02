package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "regexp"
    "text/template"
)

const lenPath = len("/view/")
const tplPath = "tpl"
const wikiPath = "wiki"

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9-]+$")
var templates = make(map[string]*template.Template)
var err error

type Page struct {
    Title string
    Body string
}

func init() {
    for _, tmpl := range []string{"edit", "view"} {
        templates[tmpl] = template.Must(template.ParseFiles(tplPath + "/" + tmpl + ".html"))
    }
}

func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
    err := http.ListenAndServe(":8088", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        title := r.URL.Path[lenPath:]
        if !titleValidator.MatchString(title) {
            http.NotFound(w, r)
            return
        }
        fn(w, r, title)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := load(title)
    if err != nil {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := load(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: body}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates[tmpl].Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func (p *Page) save() error {
    filename := wikiPath + "/" + p.Title + ".txt"
    return ioutil.WriteFile(filename, []byte(p.Body), 0755)
}

func load(title string) (*Page, error) {
    filename := wikiPath + "/" + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: string(body)}, nil
}