package monitor

import (
	"sync"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/planetmint/planetmint-go/config"
	"github.com/syndtr/goleveldb/leveldb"
)

type MQTTMonitorClientI interface {
	AddParticipant(address string, lastSeenTS int64) (err error)
	SelectPoPParticipantsOutOfActiveActors() (challenger string, challengee string, err error)
	GetActiveActorCount() (count uint64)
	Start() (err error)
}

var monitorMutex sync.RWMutex
var mqttLogger log.Logger
var mqttMonitorInstance MQTTMonitorClientI

func SetMqttMonitorInstance(monitorInstance MQTTMonitorClientI) {
	monitorMutex.Lock()
	defer monitorMutex.Unlock()
	mqttMonitorInstance = monitorInstance
}

func GetMqttMonitorInstance() (monitorInstance MQTTMonitorClientI) {
	monitorMutex.Lock()
	defer monitorMutex.Unlock()
	return mqttMonitorInstance
}

func LazyMqttMonitorLoader(logger log.Logger, homeDir string) {
	monitorMutex.RLock()
	tmpInstance := mqttMonitorInstance
	monitorMutex.RUnlock()
	if tmpInstance != nil {
		return
	}
	if logger != nil {
		mqttLogger = logger
	}
	if homeDir == "" {
		homeDir = "./"
	}
	aciveActorsDB, err := leveldb.OpenFile(homeDir+"activeActors.db", nil)
	if err != nil {
		panic(err)
	}

	SetMqttMonitorInstance(NewMqttMonitorService(aciveActorsDB, *config.GetConfig()))
}

func SelectPoPParticipantsOutOfActiveActors() (challenger string, challengee string, err error) {
	monitorMutex.RLock()
	defer monitorMutex.RUnlock()
	return mqttMonitorInstance.SelectPoPParticipantsOutOfActiveActors()
}

func Start() (err error) {
	err = mqttMonitorInstance.Start()
	return
}

func AddParticipant(address string, lastSeenTS int64) (err error) {
	monitorMutex.RLock()
	defer monitorMutex.RUnlock()
	return mqttMonitorInstance.AddParticipant(address, lastSeenTS)
}

func GetActiveActorCount() (count uint64) {
	monitorMutex.RLock()
	defer monitorMutex.RUnlock()
	return mqttMonitorInstance.GetActiveActorCount()
}

func Log(msg string) {
	if mqttLogger == nil {
		return
	}
	mqttLogger.Info("[app] [monitor] " + msg)
}
