package util

import (
	"testing"

	"github.com/planetmint/planetmint-go/util/mocks"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Use MQTT mock client
	MQTTClient = &mocks.MockMQTTClient{}
}

func TestSendMqttPopInitMessages(t *testing.T) {
	t.Parallel()
	// var challenge types.Challenge
	var challenge Challenge
	challenge.Initiator = ""
	challenge.Challengee = "plmnt15gdanx0nm2lwsx30a6wft7429p32dhzaq37c06"
	challenge.Challenger = "plmnt1683t0us0r85840nsepx6jrk2kjxw7zrcnkf0rp"
	challenge.Height = 58
	err := sendMqttPopInitMessages(challenge)
	assert.NoError(t, err)
}

func TestGetMqttStatusOfParticipantMocked(t *testing.T) {
	t.Parallel()
	participant := "plmnt15gdanx0nm2lwsx30a6wft7429p32dhzaq37c06"
	isAvailable, err := GetMqttStatusOfParticipant(participant, 200)
	assert.NoError(t, err)
	assert.True(t, isAvailable)
}

func TestToJSON(t *testing.T) {
	t.Parallel()
	payload := []byte(`{"Time":"2024-03-26T11:50:42","Uptime":"0T00:50:19","UptimeSec":3019,"Heap":97,"SleepMode":"Dynamic","Sleep":10,"LoadAvg":99,"MqttCount":2,"Berry":{"HeapUsed":27,"Objects":491},"POWER1":"ON","POWER2":"ON","Dimmer":17,"Color":"00182C","HSBColor":"207,100,17","Channel":[0,9,17],"Scheme":0,"Width":1,"Fade":"OFF","Speed":1,"LedTable":"ON","Wifi":{"AP":1,"SSId":"UPC5729E56","BSSId":"C2:14:7E:6F:BC:C5","Channel":11,"Mode":"11n","RSSI":96,"Signal":-52,"LinkCount":1,"Downtime":"0T00:00:10"}}`)

	result, err := ToJSON(payload)
	assert.NoError(t, err)
	assert.Equal(t, 21, len(result))
}
