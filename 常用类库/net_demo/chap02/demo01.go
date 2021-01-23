package chap02

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	//files := []string{"templates/layout.html", "templates/navbar.html", "templages/index.html"}
	//templates := template.Must(template.ParseFiles(files...))
	//threads, err := data.threa
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    "1231",
		HttpOnly: true,
	}
	http.SetCookie(writer, &cookie)
	http.Redirect(writer, request, "/", 302)
	_, _ = fmt.Fprintf(writer, "hello world :%s", request.URL.Path[1:])
}

func Demo01Test() {
	mux := http.NewServeMux()                     //多路复用器
	files := http.FileServer(http.Dir("/public")) //能够为指定目录中的静态文件服务的处理器
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index) //将发送到根URL的请求重定向到处理器。
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}
