package session

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClean1(t *testing.T) {

	//net.InterfaceAddrs()
	interfaces, _ := net.Interfaces()
	for _, ifac := range interfaces {
		fmt.Println(ifac.HardwareAddr)
	}
}
func TestClean2(t *testing.T) {
	name, _ := net.InterfaceByName("localhost")
	fmt.Print(name)

	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		fmt.Println(address.Network(), address.String())
	}

}
func TestRand2(t *testing.T) {
	ibuff := [16]byte{}
	rand.Read(ibuff[:])
	fmt.Print(hex.EncodeToString(ibuff[:]))

}
func TestGC(t *testing.T) {

	manager := DefaultManager()
	manager.lifeTime = 5 * 1000
	manager.StartGC()
	manager.NewSession()
	manager.NewSession()
	manager.NewSession()
	manager.NewSession()
	sessions := manager.sessions
	sessions.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		return true
	})
	time.Sleep(time.Second * 7)

	sessions.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		return true
	})

}

func TestPrintf(t *testing.T) {
	time.Now().String()
	fmt.Printf("%s", time.Now().Format("2006-01-02 15:04:05"))

}
