package mocks

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MockMQTTClient is the mock mqtt client
type MockMQTTClient struct {
	ConnectFunc    func() mqtt.Token
	DisconnectFunc func(quiesce uint)
	PublishFunc    func(topic string, qos byte, retained bool, payload interface{}) mqtt.Token
}

// GetConnectFunc fetches the mock client's `Connect` func
func GetConnectFunc() mqtt.Token {
	token := mqtt.DummyToken{}
	return &token
}

// GetDisconnectFunc fetches the mock client's `Disconnect` func
func GetDisconnectFunc(_ uint) {
}

// GetPublishFunc fetches the mock client's `Publish` func
func GetPublishFunc(_ string, _ byte, _ bool, _ interface{}) mqtt.Token {
	token := mqtt.DummyToken{}
	return &token
}

// Connect is the mock client's `Disconnect` func
func (m *MockMQTTClient) Connect() mqtt.Token {
	return GetConnectFunc()
}

// Disconnect is the mock client's `Disconnect` func
func (m *MockMQTTClient) Disconnect(quiesce uint) {
	GetDisconnectFunc(quiesce)
}

// Publish is the mock client's `Publish` func
func (m *MockMQTTClient) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return GetPublishFunc(topic, qos, retained, payload)
}
