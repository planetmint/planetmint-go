package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
)

// MockMQTTMonitorClientI is a mock of MQTTMonitorClientI interface.
type MockMQTTMonitorClientI struct {
	myStringList []string
}

// AddParticipant mocks base method.
func (m *MockMQTTMonitorClientI) AddParticipant(address string, _ int64) error {
	m.myStringList = append(m.myStringList, address)

	return nil
}

// SelectPoPParticipantsOutOfActiveActors mocks base method.
func (m *MockMQTTMonitorClientI) SelectPoPParticipantsOutOfActiveActors() (string, string, error) {
	var challenger, challengee string
	amount := len(m.myStringList)
	if amount >= 2 {
		challenger = m.myStringList[amount-2]
		challengee = m.myStringList[amount-1]
	}
	return challenger, challengee, nil
}

// SetContext mocks base method.
func (m *MockMQTTMonitorClientI) SetContext(_ types.Context) {
}

// Start mocks base method.
func (m *MockMQTTMonitorClientI) Start() error {
	return nil
}
