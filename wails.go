package main

import (
	"log"
	"sync"

	"github.com/wailsapp/wails"
)

//MyBridge -  Internal Bridge
type MyBridge struct {
	Runtime *wails.Runtime

	status string
	sync.RWMutex
	inStatus chan string
}

var internalMyBridge = MyBridge{
	status:   "Hello, World! Initial",
	inStatus: make(chan string),
}

func (mb *MyBridge) wailsRunner() {
	for {
		select {
		case newStat := <-mb.inStatus:
			mb.Lock()
			mb.status = newStat
			mb.Unlock()
			mb.Runtime.Events.Emit("update_status", newStat)
		}
	}
}

//WailsInit Инициализация структуры общения с MyBridge
func (mb *MyBridge) WailsInit(r *wails.Runtime) error {
	mb.Runtime = r
	mb.inStatus = make(chan string)
	r.Events.On("initialised", func(data ...interface{}) {
		log.Printf("Got Client!", data[0])
	})

	go mb.wailsRunner()

	return nil
}

//GetMyBridge инициализация нового экземпляра для MyBridge
func GetMyBridge() *MyBridge {

	return &internalMyBridge
}

//GetMyBridgeStatus собрали статус
func (mb *MyBridge) GetMyBridgeStatus() string {
	mb.RLock()
	defer mb.RUnlock()
	return mb.status
}

//SetStatus Установка статуса
func (mb *MyBridge) SetStatus(data string) {
	log.Printf("We Have New Status %s", data)
	mb.inStatus <- data
}
