package monitor

import (
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

var MqttMonitorInstance MQTTMonitorClientI

func LazyMqttMonitorLoader(homeDir string) {
	if MqttMonitorInstance != nil {
		return
	}
	if homeDir == "" {
		homeDir = "./"
	}
	aciveActorsDB, err := leveldb.OpenFile(homeDir+"activeActors.db", nil)
	if err != nil {
		panic(err)
	}
	MqttMonitorInstance = NewMqttMonitorService(aciveActorsDB, *config.GetConfig())
	err = MqttMonitorInstance.Start()
	if err != nil {
		panic(err)
	}
}
