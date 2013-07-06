/**
 * User: guoyao
 * Time: 07-05-2013 14:38
 * Blog: http://www.guoyao.me
 */

package session

import (
	"time"
	"fmt"
)

type Session struct {
	SessionID string
	Value string
	LastAccess time.Time
}

const (
	CookieName = "go_cookie"
	MaxLifeTime int64 = 30
)

var sessions map[string]*Session

func Add(value string) {
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
		if sessions[key].LastAccess.Unix() + MaxLifeTime < time.Now().Unix() {
			delete(sessions, key)
		}
	}
	time.AfterFunc(time.Duration(MaxLifeTime), func() {
		GC()
	})
}

func init() {
	sessions = make(map[string]*Session)
	go GC()
}
