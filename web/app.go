// 模块化、初始化时进行缓存、输入安全验证、通过匿名函数和闭包进行代码简化
// 参考 https://golang.org/doc/articles/wiki/
package web

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "./web/app/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "./web/app/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// 初始化时将模版进行缓存 // 模版名称 和 模版地址 相对应
var templates = template.Must(template.ParseFiles("./web/app/edit.html", "./web/app/view.html"))

// 渲染模版函数 // 模块化
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// 方式一
	// t, err := template.ParseFiles("./web/app/" + tmpl + ".html")
	// if err != nil {
	// 	// fmt.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// 方式二
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 文件名（标题）验证
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$") // 正则表达，分组查询

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	fmt.Println(m[0], m[1], m[2])
	return m[2], nil // The title is the second subexpression.
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// 方式一
	// title := r.URL.Path[len("/view/"):]
	// 方式二
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// 方式一
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// 方式二
	// t, _ := template.ParseFiles("./web/app/view.html")
	// t.Execute(w, p)

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// 方式一
	// title := r.URL.Path[len("/edit/"):]
	// 方式二
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// 方式一
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	//     "<form action=\"/save/%s\" method=\"POST\">"+
	//     "<textarea name=\"body\">%s</textarea><br>"+
	//     "<input type=\"submit\" value=\"Save\">"+
	//     "</form>",
	//     p.Title, p.Title, p.Body)
	// }
	// 方式二
	// t, err := template.ParseFiles("./web/app/edit.html")
	// t.Execute(w, p)

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// 方式一
	// title := r.URL.Path[len("/save/"):]
	// 方式二
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// 通过匿名函数和闭包来封装控制器 // 难点
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func AppWrite() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

func AppRun() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
