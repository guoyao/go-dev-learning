/**
 * User: guoyao
 * Time: 07-05-2013 13:59
 * Blog: http://www.guoyao.me
 */

package main

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"log"
	"session"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	fmt.Println("Http server start at:", 9090)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Hello Go Simple Blog")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("./webapp/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if username == "guoyao" && password == "123456" {
			cookie := http.Cookie{Name: session.CookieName, Value: username, Path: "/", HttpOnly: true, MaxAge: int(session.MaxLifeTime)}
			http.SetCookie(w, &cookie)
			session.Add(username, password)
			http.Redirect(w, r, "/", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}
