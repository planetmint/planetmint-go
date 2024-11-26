package monitor

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/syndtr/goleveldb/leveldb"
)

var MonitorMQTTClient util.MQTTClientI

type MqttMonitor struct {
	db                          *leveldb.DB
	dbMutex                     sync.Mutex // Mutex to synchronize write operations
	CleanupPeriodicityInMinutes time.Duration
	config                      config.Config
	numberOfElementsMutex       sync.RWMutex
	numberOfElements            int64
	isTerminated                bool
	terminationMutex            sync.RWMutex
	maxRetries                  time.Duration
	lostConnection              bool
	lostConnectionMutex         sync.Mutex
	clientMutex                 sync.Mutex
	localMqttClient             util.MQTTClientI
}

func (mms *MqttMonitor) Terminate() {
	mms.terminationMutex.Lock()
	defer mms.terminationMutex.Unlock()
	mms.isTerminated = true
}

func (mms *MqttMonitor) IsTerminated() (isTerminated bool) {
	mms.terminationMutex.RLock()
	defer mms.terminationMutex.RUnlock()
	return mms.isTerminated
}

func (mms *MqttMonitor) getNumDBElements() int64 {
	mms.numberOfElementsMutex.RLock()
	defer mms.numberOfElementsMutex.RUnlock()
	return mms.numberOfElements
}

func (mms *MqttMonitor) setNumDBElements(numElements int64) {
	mms.numberOfElementsMutex.Lock()
	defer mms.numberOfElementsMutex.Unlock()
	mms.numberOfElements = numElements
}

func getClientID() (address string) {
	conf := config.GetConfig()
	validatorAddress := conf.GetNodeAddress()
	address = "monitor-" + validatorAddress
	return
}

func (mms *MqttMonitor) lazyLoadMonitorMQTTClient() util.MQTTClientI {
	if MonitorMQTTClient != nil {
		return MonitorMQTTClient
	}

	conf := config.GetConfig()
	hostPort := net.JoinHostPort(conf.MqttDomain, strconv.FormatInt(int64(conf.MqttPort), 10))
	uri := "tcp://" + hostPort
	if conf.MqttTLS {
		uri = "ssl://" + hostPort
	}

	opts := mqtt.NewClientOptions().AddBroker(uri).SetKeepAlive(time.Second * 60).SetCleanSession(true)
	opts.SetClientID(getClientID())
	opts.SetUsername(conf.MqttUser)
	opts.SetPassword(conf.MqttPassword)
	opts.SetConnectionLostHandler(mms.onConnectionLost)
	if conf.MqttTLS {
		tlsConfig := &tls.Config{}
		opts.SetTLSConfig(tlsConfig)
	}

	Log("create new client")
	client := mqtt.NewClient(opts)
	return client
}

func NewMqttMonitorService(db *leveldb.DB, config config.Config) *MqttMonitor {
	service := &MqttMonitor{db: db, config: config, numberOfElements: 0, CleanupPeriodicityInMinutes: 10}
	return service
}

func (mms *MqttMonitor) runPeriodicTasks() {
	tickerRestablishConnection := time.NewTicker(2 * time.Minute)
	tickerCleanup := time.NewTicker(5 * time.Minute)
	defer tickerRestablishConnection.Stop()
	defer tickerCleanup.Stop()

	for !mms.IsTerminated() {
		select {
		case <-tickerRestablishConnection.C:
			go mms.MonitorActiveParticipants()
		case <-tickerCleanup.C:
			go mms.CleanupDB()
		}
	}
}

func (mms *MqttMonitor) Start() (err error) {
	amount, err := mms.getAmountOfElements()
	if err != nil {
		return
	}
	mms.setNumDBElements(amount)
	go mms.runPeriodicTasks()
	go mms.MonitorActiveParticipants()
	go mms.CleanupDB()
	return
}
func (mms *MqttMonitor) getRandomNumbers() (challenger int, challengee int) {
	for challenger == challengee {
		// Generate random numbers
		numElements := int(mms.getNumDBElements())
		challenger = rand.Intn(numElements)
		challengee = rand.Intn(numElements)
	}
	return
}
func (mms *MqttMonitor) SelectPoPParticipantsOutOfActiveActors() (challenger string, challengee string, err error) {
	numElements := int(mms.getNumDBElements())
	if numElements < 2 {
		return
	}
	randomChallenger, randomChallengee := mms.getRandomNumbers()
	Log("number of elements: " + strconv.Itoa(numElements))
	Log("selected IDs: " + strconv.Itoa(randomChallenger) + " " + strconv.Itoa(randomChallengee))
	iter := mms.db.NewIterator(nil, nil)
	defer iter.Release()
	count := 0
	found := 0
	var lastSeen LastSeenEvent
	for iter.Next() {
		if count == randomChallenger {
			lastSeen, err = mms.getDataFromIter(iter)
			if err != nil {
				Log("could not get Data from ID" + strconv.Itoa(randomChallenger))
				return
			}
			challenger = lastSeen.Address
			found++
		} else if count == randomChallengee {
			lastSeen, err = mms.getDataFromIter(iter)
			if err != nil {
				Log("could not get Data from ID" + strconv.Itoa(randomChallengee))
				return
			}
			challengee = lastSeen.Address
			found++
		}

		count++
		if found == 2 {
			break
		}
	}
	Log("challenger, challengee: " + challenger + " " + challengee)
	return
}

func (mms *MqttMonitor) GetActiveActorCount() (count uint64) {
	count = uint64(mms.getNumDBElements())
	return
}

func (mms *MqttMonitor) MqttMsgHandler(_ mqtt.Client, msg mqtt.Message) {
	if mms.IsTerminated() {
		return
	}
	topicParts := strings.Split(msg.Topic(), "/")
	if len(topicParts) != 3 {
		return
	}
	if topicParts[0] != "tele" {
		return
	}
	if topicParts[2] != "STATE" {
		return
	}
	address := topicParts[1]
	valid, err := util.IsValidAddress(address)
	if err != nil || !valid {
		return
	}

	active, err := IsLegitMachineAddress(address)
	if err != nil || !active {
		return
	}

	unixTime := time.Now().Unix()
	err = mms.AddParticipant(address, unixTime)

	if err != nil {
		Log("error adding active actor to DB: " + address + " " + err.Error())
	} else {
		Log("added active actor to DB: " + address)
	}
}

func IsLegitMachineAddress(address string) (active bool, err error) {
	url := "http://localhost:1317/planetmint/machine/address/" + address

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		Log("cannot send machine query request " + err.Error())
		return
	}

	// Set the header
	req.Header.Set("Accept", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		Log("cannot connect to server: " + err.Error())
		return
	}

	// Close the response body
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Log("cannot read response: " + err.Error())
		return
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		Log("unexpected status code: " + string(body))
		return
	}

	// Unmarshal the response body into a map
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		Log("cannot unmarshal response " + err.Error())
		return
	}

	// Check if the "info" key exists
	machineValue, ok := data["machine"]
	if !ok {
		Log("response does not contain the required machine")
		return
	}
	machineMap, ok := machineValue.(map[string]interface{})
	if !ok {
		Log("cannot convert machine map")
		return
	}
	addressMap, ok := machineMap["address"]
	if !ok {
		Log("response does not contain the required name")
		return
	}
	value, ok := addressMap.(string)
	if !ok || value != address {
		Log("return machine is not the required one")
		return
	}

	err = nil
	active = true
	return
}

func (mms *MqttMonitor) onConnectionLost(_ mqtt.Client, err error) {
	Log("connection lost: " + err.Error())
	// Handle connection loss here (e.g., reconnect attempts, logging)
	if !mms.IsTerminated() {
		mms.setConnectionStatus(true)
	}
}

func (mms *MqttMonitor) MonitorActiveParticipants() {
	mqttClient, err := mms.initializeClient()
	if err != nil {
		Log(err.Error())
		return
	}

	// Maximum reconnection attempts (adjust as needed)
	mms.SetMaxRetries()

	for !mms.IsTerminated() && mms.maxRetries > 0 {
		if !mms.connectClient(mqttClient) {
			continue
		}

		if !mms.subscribeToTopic(mqttClient) {
			continue
		}

		mms.monitorConnection(mqttClient)
	}

	mms.handleConnectionTermination()
}

func (mms *MqttMonitor) initializeClient() (mqttClient util.MQTTClientI, err error) {
	mms.clientMutex.Lock()
	defer mms.clientMutex.Unlock()

	if mms.localMqttClient != nil {
		return nil, errors.New("client is still working")
	}

	mms.localMqttClient = mms.lazyLoadMonitorMQTTClient()
	return mms.localMqttClient, nil
}

func (mms *MqttMonitor) connectClient(mqttClient util.MQTTClientI) bool {
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		Log("error connecting to mqtt: " + token.Error().Error())
		mms.maxRetries--
		time.Sleep(time.Second * 5)
		return false
	}

	mms.setConnectionStatus(false)
	Log("established connection")
	return true
}

func (mms *MqttMonitor) subscribeToTopic(mqttClient util.MQTTClientI) bool {
	messageHandler := mqtt.MessageHandler(mms.MqttMsgHandler)
	subscriptionTopic := "tele/#"

	if token := mqttClient.Subscribe(subscriptionTopic, 0, messageHandler); token.Wait() && token.Error() != nil {
		Log("error registering the mqtt subscription: " + token.Error().Error())
		return false
	}

	Log("subscribed to tele/# channels")
	return true
}

func (mms *MqttMonitor) monitorConnection(mqttClient util.MQTTClientI) {
	for !mms.IsTerminated() {
		if mms.isConnectionLost(mqttClient) {
			Log("retry establishing a connection")
			break
		}

		SendUpdateMessage(mqttClient)
		mms.SetMaxRetries()
		time.Sleep(60 * time.Second)
	}
}

func (mms *MqttMonitor) isConnectionLost(mqttClient util.MQTTClientI) bool {
	mms.lostConnectionMutex.Lock()
	defer mms.lostConnectionMutex.Unlock()

	return !mqttClient.IsConnected() || !mqttClient.IsConnectionOpen() || mms.lostConnection
}

func (mms *MqttMonitor) setConnectionStatus(lost bool) {
	mms.lostConnectionMutex.Lock()
	defer mms.lostConnectionMutex.Unlock()
	mms.lostConnection = lost
}

func (mms *MqttMonitor) handleConnectionTermination() {
	if mms.maxRetries == 0 {
		Log("reached maximum reconnection attempts. Exiting. New client will be activated soon.")
	}

	mms.clientMutex.Lock()
	defer mms.clientMutex.Unlock()
	mms.localMqttClient = nil
}

func SendUpdateMessage(mqttClient util.MQTTClientI) {
	// Publish message
	now := time.Now().Format("2006-01-02 15:04:05") // Adjust format as needed
	token := mqttClient.Publish("tele/"+getClientID(), 1, false, now)
	token.Wait()
}

func (mms *MqttMonitor) SetMaxRetries() {
	mms.maxRetries = 5
}
