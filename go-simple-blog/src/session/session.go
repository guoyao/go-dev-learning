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

type Manager struct {
	CookieName string
	MaxLifeTime int64
	sessions map[string]*Session
}

var SessionManager Manager;

func (manager *Manager) AddSession(cookieValue string) {
	if _, ok := manager.sessions[cookieValue]; ok {
		fmt.Printf("Add new session failed, session '%s' already exists.", cookieValue)
	} else {
		manager.sessions[cookieValue] = &Session{cookieValue, cookieValue, time.Now()}
	}
}

func (manager *Manager) exists(key string) bool {
	_, ok := manager.sessions[key]
	return ok
}

func (manager *Manager) UpdateSession(cookieValue string) {
	if manager.exists(cookieValue) {
		manager.sessions[cookieValue].LastAccess = time.Now()
	}
}

func (manager *Manager) gc() {
	for key, _ := range manager.sessions {
		if manager.sessions[key].LastAccess.Unix() + manager.MaxLifeTime < time.Now().Unix() {
			delete(manager.sessions, key)
		}
	}
	time.AfterFunc(time.Duration(manager.MaxLifeTime), func() {
		manager.gc()
	})
}

func gc() {
	SessionManager.gc()
}

func init() {
	SessionManager = Manager{"GO_SESSION", 20, make(map[string]*Session)}
	go gc()
}
