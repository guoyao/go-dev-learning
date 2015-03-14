/**
 * User: guoyao
 * Time: 07-01-2013 14:24
 * Blog: http://www.guoyao.me
 */

package main

import (
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
