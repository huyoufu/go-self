package session

import "time"

var SessionCookieName = "jksessionid"

type Session interface {
	Id() string
	Get(name string) interface{}
	Set(name string, value interface{})
	Remove(name string)
	Invalidate()
	LastAccessedTime() time.Time
	CreationTime() time.Time
	Names() []string
	access()
	IsNew() bool
}
