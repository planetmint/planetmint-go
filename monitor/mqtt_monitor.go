package monitor

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/syndtr/goleveldb/leveldb"
)

var MonitorMQTTClient util.MQTTClientI
var clientMutex sync.Mutex
var localMqttClient util.MQTTClientI

type MqttMonitor struct {
	db                          *leveldb.DB
	dbMutex                     sync.Mutex // Mutex to synchronize write operations
	CleanupPeriodicityInMinutes time.Duration
	config                      config.Config
	numberOfElements            int64
	sdkContext                  *sdk.Context
	contextMutex                sync.Mutex
	isTerminated                bool
	terminationMutex            sync.Mutex
	maxRetries                  time.Duration
	lostConnection              bool
	lostConnectionMutex         sync.Mutex
}

func (mms *MqttMonitor) Terminate() {
	mms.terminationMutex.Lock()
	mms.isTerminated = true
	mms.terminationMutex.Unlock()
}

func (mms *MqttMonitor) IsTerminated() (isTerminated bool) {
	mms.terminationMutex.Lock()
	isTerminated = mms.isTerminated
	mms.terminationMutex.Unlock()
	return
}

func getClientID() string {
	conf := config.GetConfig()
	return "monitor-" + conf.ValidatorAddress
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

	log.Println("[app] [Monitor] create new client")
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
	mms.numberOfElements = amount
	go mms.runPeriodicTasks()
	go mms.MonitorActiveParticipants()
	go mms.CleanupDB()
	return
}
func (mms *MqttMonitor) getRandomNumbers() (challenger int, challengee int) {
	for challenger == challengee {
		// Generate random numbers
		challenger = rand.Intn(int(mms.numberOfElements))
		challengee = rand.Intn(int(mms.numberOfElements))
	}
	return
}
func (mms *MqttMonitor) SelectPoPParticipantsOutOfActiveActors() (challenger string, challengee string, err error) {
	if mms.numberOfElements < 2 {
		return
	}
	randomChallenger, randomChallengee := mms.getRandomNumbers()
	log.Println("[app] [Monitor] number of elements: " + strconv.Itoa(int(mms.numberOfElements)))
	log.Println("[app] [Monitor] selected IDs: " + strconv.Itoa(randomChallenger) + " " + strconv.Itoa(randomChallengee))
	iter := mms.db.NewIterator(nil, nil)
	defer iter.Release()
	count := 0
	found := 0
	var lastSeen LastSeenEvent
	for iter.Next() {
		if count == randomChallenger {
			lastSeen, err = mms.getDataFromIter(iter)
			if err != nil {
				log.Println("[app] [Monitor] could not get Data from ID" + strconv.Itoa(randomChallenger))
				return
			}
			challenger = lastSeen.Address
			found++
		} else if count == randomChallengee {
			lastSeen, err = mms.getDataFromIter(iter)
			if err != nil {
				log.Println("[app] [Monitor] could not get Data from ID" + strconv.Itoa(randomChallengee))
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
	log.Println("[app] [Monitor] challenger, challengee: " + challenger + " " + challengee)
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

	unixTime := time.Now().Unix()
	err = mms.AddParticipant(address, unixTime)

	if err != nil {
		log.Println("[app] [Monitor] error adding active actor to DB: " + address + " " + err.Error())
	} else {
		log.Println("[app] [Monitor] added active actor to DB: " + address)
	}
}

func (mms *MqttMonitor) onConnectionLost(_ mqtt.Client, err error) {
	log.Println("[app] [Monitor] Connection lost: " + err.Error())
	// Handle connection loss here (e.g., reconnect attempts, logging)
	if !mms.IsTerminated() {
		mms.lostConnectionMutex.Lock()
		mms.lostConnection = true
		mms.lostConnectionMutex.Unlock()
	}
}

func (mms *MqttMonitor) MonitorActiveParticipants() {
	clientMutex.Lock()
	if localMqttClient != nil {
		log.Println("[app] [Monitor] client is still working")
		clientMutex.Unlock()
		return
	}
	localMqttClient = mms.lazyLoadMonitorMQTTClient()
	mqttClient := localMqttClient
	clientMutex.Unlock()

	// Maximum reconnection attempts (adjust as needed)
	mms.SetMaxRetries()
	for !mms.IsTerminated() && mms.maxRetries > 0 {
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			log.Println("[app] [Monitor] error connecting to mqtt: " + token.Error().Error())
			mms.maxRetries--
			time.Sleep(time.Second * 5)
			continue
		}
		mms.lostConnectionMutex.Lock()
		mms.lostConnection = false
		mms.lostConnectionMutex.Unlock()

		log.Println("[app] [Monitor] established connection")

		var messageHandler mqtt.MessageHandler = mms.MqttMsgHandler

		// Subscribe to a topic
		subscriptionTopic := "tele/#"
		if token := mqttClient.Subscribe(subscriptionTopic, 0, messageHandler); token.Wait() && token.Error() != nil {
			log.Println("[app] [Monitor] error registering the mqtt subscription: " + token.Error().Error())
			continue
		}
		log.Println("[app] [Monitor] subscribed to tele/# channels")

		for !mms.IsTerminated() {
			mms.lostConnectionMutex.Lock()
			lostConnectionEvent := mms.lostConnection
			mms.lostConnectionMutex.Unlock()
			if !mqttClient.IsConnected() || !mqttClient.IsConnectionOpen() || lostConnectionEvent {
				log.Println("[app] [Monitor] retry establishing a connection")
				break // Exit inner loop on disconnect
			}

			SendUpdateMessage(mqttClient)
			mms.SetMaxRetries()
			time.Sleep(60 * time.Second) // Adjust sleep time based on your needs
		}
	}

	if mms.maxRetries == 0 {
		log.Println("[app] [Monitor] Reached maximum reconnection attempts. Exiting. New client will be activated soon.")
	}

	clientMutex.Lock()
	localMqttClient = nil
	clientMutex.Unlock()
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

func (mms *MqttMonitor) Log(msg string) {
	mms.contextMutex.Lock()
	localContext := mms.sdkContext
	mms.contextMutex.Unlock()
	if localContext != nil {
		util.GetAppLogger().Info(*localContext, msg)
	}
}
