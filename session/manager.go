package session

import (
	"fmt"
	. "github.com/huyoufu/go-self/logger"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type providerType uint8

const (
	runtime providerType = iota // default
	redis                       //redis实现
)

type Manager struct {
	sessions sync.Map //包含的所有的session
	lifeTime int64    //活跃时间 以毫秒为 单位
	pType    providerType
}

func DefaultManager() *Manager {
	return &Manager{
		sessions: sync.Map{},
		lifeTime: 30 * 60 * 1000,
		/*lifeTime:5*1000,*/
		pType: runtime,
	}
}
func (m *Manager) NewSession() Session {

	sessionId := uuid.Must(uuid.NewV4()).String()
	session := m.newSession(sessionId)
	m.sessions.Store(sessionId, session)
	return session
}

func (m *Manager) newSession(sessionId string) Session {
	switch m.pType {
	case runtime:
		return newRuntimeSession(sessionId, m)
	case redis:
		fmt.Println("暂未支持")
		return nil
	default:
		panic("不支持的类型")
	}
}

func (m *Manager) GetSession(sessionId string) (s Session) {
	s = nil
	m.sessions.Range(func(key, value interface{}) bool {
		if sessionId == key {
			s = value.(Session)
			return false
		} else {
			s = nil
			return true
		}
	})
	if s == nil {
		//如果为空 就重新创建一个
		s = m.newSession(sessionId)
		m.sessions.Store(sessionId, s)
	}

	//重新存活时间
	s.access()

	return
}
func Access(session Session) {
	session.access()
}
func (m *Manager) Remove(session Session) {
	m.sessions.Delete(session.Id())
}
func (m *Manager) StartGC() {
	m.gc()
}
func (m *Manager) gc() {
	Log.Debugf("start sessions gc")
	m.sessions.Range(func(key, value interface{}) bool {
		s := value.(Session)
		c1 := time.Now().UnixNano() / 1e6
		c2 := s.LastAccessedTime().UnixNano() / 1e6
		if c1-c2 > m.lifeTime {
			Log.Debugf("正在删除:%d", s.Id())
			m.sessions.Delete(key)
		}
		return true
	})
	time.AfterFunc(time.Minute, m.gc)
}
