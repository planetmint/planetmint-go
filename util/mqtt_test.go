package util

import (
	"testing"

	"github.com/planetmint/planetmint-go/util/mocks"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Use MQTT mock client
	MQTTClient = &mocks.MockMQTTClient{}
}

func TestSendMqttPopInitMessages(t *testing.T) {
	t.Parallel()
	var challenge types.Challenge
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
	isAvailable, err := GetMqttStatusOfParticipant(participant)
	assert.NoError(t, err)
	assert.True(t, isAvailable)
}
