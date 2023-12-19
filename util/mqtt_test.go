package util

import (
	"testing"

	"github.com/planetmint/planetmint-go/x/dao/types"
)

func TestSendMqttMessages(t *testing.T) {
	t.Parallel()
	t.Skip("Skip this test case as this test case expects a working MQTT connection" +
		"the test case is intended to work manually.")

	var challenge types.Challenge
	challenge.Initiator = ""
	challenge.Challengee = "plmnt15gdanx0nm2lwsx30a6wft7429p32dhzaq37c06"
	challenge.Challenger = "plmnt1683t0us0r85840nsepx6jrk2kjxw7zrcnkf0rp"
	challenge.Height = 58
	sendMqttMessages(challenge)
}
