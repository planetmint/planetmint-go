package mocks

import "log"

// MockMQTTMonitorClientI is a mock of MQTTMonitorClientI interface.
type MockMQTTMonitorClientI struct {
	myStringList []string
}

// AddParticipant mocks base method.
func (m *MockMQTTMonitorClientI) AddParticipant(address string, _ int64) error {
	log.Println("[app] [Monitor] [Mock] added participant: " + address)
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
	log.Println("[app] [Monitor] [Mock] participants: " + challenger + ", " + challengee)
	return challenger, challengee, nil
}

func (m *MockMQTTMonitorClientI) GetActiveActorCount() (count uint64) {
	return uint64(len(m.myStringList))
}

// Start mocks base method.
func (m *MockMQTTMonitorClientI) Start() error {
	return nil
}
