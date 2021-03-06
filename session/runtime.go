package session

import (
	"sync"
	"time"
)

type RuntimeSession struct {
	id               string
	attributes       sync.Map
	lastAccessedTime time.Time
	creationTime     time.Time
	names            []string
	lock             *sync.RWMutex
	manager          *Manager
	isNew            bool
}

func (s *RuntimeSession) Id() string {
	return s.id
}
func (s *RuntimeSession) Get(name string) (value interface{}) {
	value, _ = s.attributes.Load(name)
	return
}
func (s *RuntimeSession) Set(name string, value interface{}) {
	s.attributes.Store(name, value)
}
func (s *RuntimeSession) Remove(name string) {
	s.attributes.Delete(name)
}
func (s *RuntimeSession) Invalidate() {
	s.manager.Remove(s)
}
func (s *RuntimeSession) LastAccessedTime() time.Time {
	return s.lastAccessedTime
}
func (s *RuntimeSession) CreationTime() time.Time {
	return s.creationTime
}
func (s *RuntimeSession) Names() []string {
	return s.names
}
func (s *RuntimeSession) access() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.lastAccessedTime = time.Now()
	s.isNew = false

}
func (s *RuntimeSession) IsNew() bool {
	return s.isNew
}

func newRuntimeSession(sessionId string, m *Manager) Session {
	return &RuntimeSession{
		id:               sessionId,
		attributes:       sync.Map{},
		lastAccessedTime: time.Now(),
		creationTime:     time.Now(),
		names:            []string{},
		lock:             &sync.RWMutex{},
		manager:          m,
		isNew:            true,
	}
}
