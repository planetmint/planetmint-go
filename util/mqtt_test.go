package util

import (
	"testing"

	"github.com/planetmint/planetmint-go/util/mocks"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/assert"
)

func TestSendMqttMessages(t *testing.T) {
	t.Parallel()
	var challenge types.Challenge
	challenge.Initiator = ""
	challenge.Challengee = "plmnt15gdanx0nm2lwsx30a6wft7429p32dhzaq37c06"
	challenge.Challenger = "plmnt1683t0us0r85840nsepx6jrk2kjxw7zrcnkf0rp"
	challenge.Height = 58
	// Use MQTT mock client
	MQTTClient = &mocks.MockMQTTClient{}
	err := sendMqttMessages(challenge)
	assert.NoError(t, err)
}
