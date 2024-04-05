package monitor

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/syndtr/goleveldb/leveldb"
)

type MQTTMonitorClientI interface {
	AddParticipant(address string, lastSeenTS int64) (err error)
	SelectPoPParticipantsOutOfActiveActors() (challenger string, challengee string, err error)
	SetContext(ctx sdk.Context)
	Start() (err error)
}

var monitorMutex sync.Mutex
var mqttMonitorInstance MQTTMonitorClientI

func SetMqttMonitorInstance(monitorInstance MQTTMonitorClientI) {
	monitorMutex.Lock()
	mqttMonitorInstance = monitorInstance
	monitorMutex.Unlock()
}

func LazyMqttMonitorLoader(homeDir string) {
	monitorMutex.Lock()
	tmpInstance := mqttMonitorInstance
	monitorMutex.Unlock()
	if tmpInstance != nil {
		return
	}
	if homeDir == "" {
		homeDir = "./"
	}
	aciveActorsDB, err := leveldb.OpenFile(homeDir+"activeActors.db", nil)
	if err != nil {
		panic(err)
	}
	monitorMutex.Lock()
	mqttMonitorInstance = NewMqttMonitorService(aciveActorsDB, *config.GetConfig())
	monitorMutex.Unlock()
	err = mqttMonitorInstance.Start()
	if err != nil {
		panic(err)
	}

}

func SetContext(ctx sdk.Context) {
	monitorMutex.Lock()
	mqttMonitorInstance.SetContext(ctx)
	monitorMutex.Unlock()
}

func SelectPoPParticipantsOutOfActiveActors() (challenger string, challengee string, err error) {
	monitorMutex.Lock()
	challenger, challengee, err = mqttMonitorInstance.SelectPoPParticipantsOutOfActiveActors()
	monitorMutex.Unlock()
	return
}

func Start() (err error) {
	err = mqttMonitorInstance.Start()
	return
}

func AddParticipant(address string, lastSeenTS int64) (err error) {
	monitorMutex.Lock()
	err = mqttMonitorInstance.AddParticipant(address, lastSeenTS)
	monitorMutex.Unlock()
	return
}
