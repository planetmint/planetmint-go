package monitor_test

import (
	"testing"
	"time"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/monitor"
	"github.com/planetmint/planetmint-go/util/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

func init() {
	// Use MQTT mock client
	monitor.MonitorMQTTClient = &mocks.MockMQTTClient{}
}

const (
	challengerInput = "plmnt1fx3x6u8k5q8kjl7pamsuwjtut8nkks8dk92dek"
	challengeeInput = "plmnt1fsaljz3xqf6vchkjxfzfrd30cdp3j4vqh298pr"
)

func TestGMonitorActiveParticipants(t *testing.T) {
	cfg := config.GetConfig()
	db, err := leveldb.Open(storage.NewMemStorage(), nil)
	assert.NoError(t, err)
	defer db.Close()

	mqttMonitor := monitor.NewMqttMonitorService(db, *cfg)
	err = mqttMonitor.Start()
	assert.NoError(t, err)

	currentTime := time.Now()
	unixTime := currentTime.Unix()
	err = mqttMonitor.AddParticipant(challengerInput, unixTime)
	assert.NoError(t, err)
	err = mqttMonitor.AddParticipant(challengeeInput, unixTime)
	assert.NoError(t, err)
	mqttMonitor.CleanupDB()

	challenger, challengee, err := mqttMonitor.SelectPoPParticipantsOutOfActiveActors()
	assert.NoError(t, err)
	assert.Contains(t, challenger, "plmnt")
	assert.Contains(t, challengee, "plmnt")
	mqttMonitor.Terminate()
}

func TestCleanupRemoval(t *testing.T) {
	cfg := config.GetConfig()
	db, err := leveldb.Open(storage.NewMemStorage(), nil)
	assert.NoError(t, err)
	defer db.Close()

	mqttMonitor := monitor.NewMqttMonitorService(db, *cfg)
	err = mqttMonitor.Start()
	assert.NoError(t, err)

	currentTime := time.Now()
	CleanupPeriodicityAgo := currentTime.Add(-1 * mqttMonitor.CleanupPeriodicityInMinutes * time.Minute)
	unixTimeNow := currentTime.Unix()
	err = mqttMonitor.AddParticipant(challengerInput, unixTimeNow)
	assert.NoError(t, err)
	err = mqttMonitor.AddParticipant(challengeeInput, CleanupPeriodicityAgo.Unix()-1)
	assert.NoError(t, err)
	mqttMonitor.CleanupDB()

	challenger, challengee, err := mqttMonitor.SelectPoPParticipantsOutOfActiveActors()
	assert.NoError(t, err)
	assert.Equal(t, "", challenger)
	assert.Contains(t, "", challengee)
	mqttMonitor.Terminate()
}

func TestCleanupPrecisionTest(t *testing.T) {
	cfg := config.GetConfig()
	db, err := leveldb.Open(storage.NewMemStorage(), nil)
	assert.NoError(t, err)
	defer db.Close()

	mqttMonitor := monitor.NewMqttMonitorService(db, *cfg)
	err = mqttMonitor.Start()
	assert.NoError(t, err)

	currentTime := time.Now()
	CleanupThresholdAgo := currentTime.Add(-1 * mqttMonitor.CleanupPeriodicityInMinutes * time.Minute)
	aboveThreshold := CleanupThresholdAgo.Unix() + 10
	unixTimeNow := currentTime.Unix()
	err = mqttMonitor.AddParticipant(challengerInput, unixTimeNow)
	assert.NoError(t, err)
	err = mqttMonitor.AddParticipant(challengeeInput, aboveThreshold)
	assert.NoError(t, err)
	mqttMonitor.CleanupDB()

	challenger, challengee, err := mqttMonitor.SelectPoPParticipantsOutOfActiveActors()
	assert.NoError(t, err)
	assert.Contains(t, challenger, "plmnt")
	assert.Contains(t, challengee, "plmnt")
	mqttMonitor.Terminate()
}

func TestIsLegitMachineAddress(t *testing.T) {
	t.SkipNow()
	active, err := monitor.IsLegitMachineAddress("plmnt1z6xmwqfnn9mvean9gsd57segawgjykpxw8hq5t")
	assert.NoError(t, err)
	assert.Equal(t, active, true)
}
