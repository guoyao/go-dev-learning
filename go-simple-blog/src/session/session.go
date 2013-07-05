/**
 * User: guoyao
 * Time: 07-05-2013 14:38
 * Blog: http://www.guoyao.me
 */

package session

import (
	"time"
	"fmt"
	"net/http"
)

type Session struct {
	SessionID string
	Value string
	LastAccess time.Time
}

const (
	cookieName = "go_cookie"
	maxLifeTime int64 = 5000
)

var sessions map[string]*Session

func Add(w http.ResponseWriter, r *http.Request, value string) {
	if _, ok := sessions[value]; ok {
		fmt.Printf("Add new session failed, session '%s' already exists.", value)
	} else {
		session := Session{value, value, time.Now()}
		sessions[value] = &session
	}
}

func Exists(key string) bool {
	_, ok := sessions[key]
	return ok
}

func Update(key string) {
	if Exists(key) {
		sessions[key].LastAccess = time.Now()
	}
}

func GC() {
	for key, _ := range sessions {
		if sessions[key].LastAccess.Unix() + maxLifeTime < time.Now().Unix() {
			delete(sessions, key)
		}
	}
}

func init() {
	sessions = make(map[string]*Session)
}
