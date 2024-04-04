package monitor_test

import (
	"testing"
	"time"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/monitor"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/util/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
)

func init() {
	// Use MQTT mock client
	util.MQTTClient = &mocks.MockMQTTClient{}
}

const (
	challengerInput = "plmnt1fx3x6u8k5q8kjl7pamsuwjtut8nkks8dk92dek"
	challengeeInput = "plmnt1fsaljz3xqf6vchkjxfzfrd30cdp3j4vqh298pr"
)

func TestGMonitorActiveParticipants(t *testing.T) {
	util.LazyLoadMQTTClient()
	cfg := config.GetConfig()
	db, err := leveldb.OpenFile("./activeActors.db", nil)
	assert.NoError(t, err)
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
}

func TestCleanupRemoval(t *testing.T) {
	util.LazyLoadMQTTClient()

	cfg := config.GetConfig()
	db, err := leveldb.OpenFile("./activeActors.db", nil)
	assert.NoError(t, err)
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
}

func TestCleanupPrecisionTest(t *testing.T) {
	util.LazyLoadMQTTClient()

	cfg := config.GetConfig()
	db, err := leveldb.OpenFile("./activeActors.db", nil)
	assert.NoError(t, err)
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
}
