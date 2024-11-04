package mocks

import (
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	callLog  []Call
	logMutex sync.Mutex
)

type Call struct {
	FuncName string
	Params   []interface{}
}

// MockMQTTClient is the mock mqtt client
type MockMQTTClient struct {
	ConnectFunc          func() mqtt.Token
	DisconnectFunc       func(quiesce uint)
	PublishFunc          func(topic string, qos byte, retained bool, payload interface{}) mqtt.Token
	SubscribeFunc        func(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token
	UnsubscribeFunc      func(topics ...string) mqtt.Token
	IsConnectedFunc      func() bool
	IsConnectionOpenFunc func() bool
	connected            bool
	connectedMutex       sync.Mutex
}

// GetConnectFunc fetches the mock client's `Connect` func
func GetConnectFunc() mqtt.Token {
	token := mqtt.DummyToken{}
	return &token
}

// GetDisconnectFunc fetches the mock client's `Disconnect` func
func GetDisconnectFunc(_ uint) {
	// not implemented at this point in time
}

// GetPublishFunc fetches the mock client's `Publish` func
func GetPublishFunc(_ string, _ byte, _ bool, _ interface{}) mqtt.Token {
	token := mqtt.DummyToken{}
	return &token
}

type message struct {
	duplicate bool
	qos       byte
	retained  bool
	topic     string
	messageID uint16
	payload   []byte
	ack       func()
	once      sync.Once
}

func (m *message) Duplicate() bool {
	return m.duplicate
}

func (m *message) Qos() byte {
	return m.qos
}

func (m *message) Retained() bool {
	return m.retained
}

func (m *message) Topic() string {
	return m.topic
}

func (m *message) MessageID() uint16 {
	return m.messageID
}

func (m *message) Payload() []byte {
	return m.payload
}

func (m *message) Ack() {
	m.once.Do(m.ack)
}

func SendCallbackResponse(topic string, callback mqtt.MessageHandler) {
	time.Sleep(100 * time.Millisecond)
	var opt mqtt.ClientOptions
	client := mqtt.NewClient(&opt)
	var msg message
	msg.topic = topic
	var msgInter mqtt.Message = &msg

	callback(client, msgInter)
}

func GetSubscribeFunc(topic string, _ byte, callback mqtt.MessageHandler) mqtt.Token {
	go SendCallbackResponse(topic, callback)
	token := mqtt.DummyToken{}
	return &token
}

func GetUnsubscribeFunc(_ ...string) mqtt.Token {
	token := mqtt.DummyToken{}
	return &token
}

// Connect is the mock client's `Disconnect` func
func (m *MockMQTTClient) Connect() mqtt.Token {
	m.logCall("Connect")
	m.connectedMutex.Lock()
	m.connected = true
	m.connectedMutex.Unlock()
	return GetConnectFunc()
}

// Disconnect is the mock client's `Disconnect` func
func (m *MockMQTTClient) Disconnect(quiesce uint) {
	m.logCall("Disconnect")
	GetDisconnectFunc(quiesce)
}

// Publish is the mock client's `Publish` func
func (m *MockMQTTClient) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	m.logCall("Publish", topic, qos, retained, payload)
	return GetPublishFunc(topic, qos, retained, payload)
}

func (m *MockMQTTClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	m.logCall("Subscribe", topic, qos, callback)
	return GetSubscribeFunc(topic, qos, callback)
}

func (m *MockMQTTClient) Unsubscribe(topics ...string) mqtt.Token {
	m.logCall("Unsubscribe", topics)
	return GetUnsubscribeFunc(topics...)
}

func (m *MockMQTTClient) IsConnected() bool {
	m.logCall("IsConnected")
	m.connectedMutex.Lock()
	connected := m.connected
	m.connectedMutex.Unlock()
	return connected
}

func (m *MockMQTTClient) IsConnectionOpen() bool {
	m.logCall("IsConnectedOpen")
	m.connectedMutex.Lock()
	connected := m.connected
	m.connectedMutex.Unlock()
	return connected
}

func GetCallLog() []Call {
	logMutex.Lock()
	defer logMutex.Unlock()
	return callLog
}

func (m *MockMQTTClient) logCall(funcName string, params ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	callLog = append(callLog, Call{FuncName: funcName, Params: params})
}
