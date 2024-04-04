package monitor

import (
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

type MqttMonitor struct {
	db                          *leveldb.DB
	dbMutex                     sync.Mutex // Mutex to synchronize write operations
	ticker                      *time.Ticker
	CleanupPeriodicityInMinutes time.Duration
	config                      config.Config
	numberOfElements            int64
	sdkContext                  *sdk.Context
	contextMutex                sync.Mutex
}

func LazyLoadMonitorMQTTClient() {
	if MonitorMQTTClient != nil {
		return
	}

	conf := config.GetConfig()
	hostPort := net.JoinHostPort(conf.MqttDomain, strconv.FormatInt(int64(conf.MqttPort), 10))
	uri := "tcp://" + hostPort

	opts := mqtt.NewClientOptions().AddBroker(uri)
	opts.SetClientID(conf.ValidatorAddress + "-monitor")
	opts.SetUsername(conf.MqttUser)
	opts.SetPassword(conf.MqttPassword)
	MonitorMQTTClient = mqtt.NewClient(opts)
}

func NewMqttMonitorService(db *leveldb.DB, config config.Config) *MqttMonitor {
	service := &MqttMonitor{db: db, config: config, numberOfElements: 0, CleanupPeriodicityInMinutes: 10}
	return service
}

func (mms *MqttMonitor) registerPeriodicTasks() {
	mms.ticker = time.NewTicker(mms.CleanupPeriodicityInMinutes * time.Minute)
	go func() {
		for range mms.ticker.C { // Loop over the ticker channel
			go mms.CleanupDB()
		}
	}()
}

func (mms *MqttMonitor) Start() (err error) {
	amount, err := mms.getAmountOfElements()
	if err != nil {
		return
	}
	mms.numberOfElements = amount
	mms.registerPeriodicTasks()
	go mms.MonitorActiveParticipants()
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
	mms.Log("[Monitor] number of elements: " + strconv.Itoa(int(mms.numberOfElements)))
	mms.Log("[Monitor] selected IDs: " + strconv.Itoa(randomChallenger) + " " + strconv.Itoa(randomChallengee))
	iter := mms.db.NewIterator(nil, nil)
	defer iter.Release()
	count := 0
	found := 0
	var lastSeen LastSeenEvent
	for iter.Next() {
		mms.Log("[Monitor] count: " + strconv.Itoa(count))
		if count == randomChallenger {
			lastSeen, err = mms.getDataFromIter(iter)
			if err != nil {
				return
			}
			challenger = lastSeen.Address
			found++
		} else if count == randomChallengee {
			lastSeen, err = mms.getDataFromIter(iter)
			if err != nil {
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
	return
}

func (mms *MqttMonitor) MqttMsgHandler(_ mqtt.Client, msg mqtt.Message) {
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
	payload, err := util.ToJSON(msg.Payload())
	if err != nil {
		return
	}

	timeString, ok := payload["Time"].(string)
	if !ok {
		return
	}
	unixTime, err := util.String2UnixTime(timeString)
	if err != nil {
		return
	}
	err = mms.AddParticipant(address, unixTime)
	if err != nil {
		mms.Log("[Monitor] error adding active actor to DB: " + address + " " + err.Error())
	} else {
		mms.Log("[Monitor] added active actor to DB: " + address)
	}
}

func (mms *MqttMonitor) MonitorActiveParticipants() {
	LazyLoadMonitorMQTTClient()
	for {
		if !MonitorMQTTClient.IsConnected() {
			if token := MonitorMQTTClient.Connect(); token.Wait() && token.Error() != nil {
				mms.Log("[Monitor] error connecting to mqtt: " + token.Error().Error())
				panic(token.Error())
			}

			var messageHandler mqtt.MessageHandler = mms.MqttMsgHandler

			// Subscribe to a topic
			subscriptionTopic := "tele/#"
			if token := MonitorMQTTClient.Subscribe(subscriptionTopic, 0, messageHandler); token.Wait() && token.Error() != nil {
				mms.Log("[Monitor] error registering the mqtt subscription: " + token.Error().Error())
				panic(token.Error())
			}
		}
		time.Sleep(30 * time.Second)
	}
}

func (mms *MqttMonitor) Log(msg string) {
	mms.contextMutex.Lock()
	if mms.sdkContext != nil {
		util.GetAppLogger().Info(*mms.sdkContext, msg)
	}
	mms.contextMutex.Unlock()
}

func (mms *MqttMonitor) SetContext(ctx sdk.Context) {
	mms.contextMutex.Lock()
	mms.sdkContext = &ctx
	mms.contextMutex.Unlock()
}
