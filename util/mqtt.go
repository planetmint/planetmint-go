package util

import (
	"crypto/tls"
	"encoding/json"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// MQTTClientI interface
type MQTTClientI interface {
	Connect() mqtt.Token
	Disconnect(quiesce uint)
	Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token
	Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token
	Unsubscribe(topics ...string) mqtt.Token
	IsConnected() bool
	IsConnectionOpen() bool
}

var (
	MQTTClient                              MQTTClientI
	mqttMachineByAddressAvailabilityMapping map[string]bool
	rwMu                                    sync.RWMutex
)

const (
	MqttCmdPrefix = "cmnd/"
)

func LazyLoadMQTTClient() {
	if MQTTClient != nil {
		return
	}

	conf := config.GetConfig()
	hostPort := net.JoinHostPort(conf.MqttDomain, strconv.FormatInt(int64(conf.MqttPort), 10))
	uri := "tcp://" + hostPort
	if conf.MqttTLS {
		uri = "ssl://" + hostPort
	}

	opts := mqtt.NewClientOptions().AddBroker(uri)
	opts.SetClientID(conf.GetNodeAddress())
	opts.SetUsername(conf.MqttUser)
	opts.SetPassword(conf.MqttPassword)
	if conf.MqttTLS {
		tlsConfig := &tls.Config{}
		opts.SetTLSConfig(tlsConfig)
	}

	MQTTClient = mqtt.NewClient(opts)
}

func init() {
	mqttMachineByAddressAvailabilityMapping = make(map[string]bool)
}

func SendMqttPopInitMessagesToServer(ctx sdk.Context, challenge types.Challenge) {
	// PoP can only be executed if at least two actors are available.
	if challenge.Challenger == "" || challenge.Challengee == "" {
		return
	}
	err := sendMqttPopInitMessages(challenge)
	if err != nil {
		GetAppLogger().Error(ctx, err, "MQTT")
		return
	}
	GetAppLogger().Info(ctx, "MQTT message successfully sent: "+challenge.String())
}

func sendMqttPopInitMessages(challenge types.Challenge) (err error) {
	LazyLoadMQTTClient()
	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}

	blockHeight := strconv.FormatInt(challenge.GetHeight(), 10)
	token := MQTTClient.Publish(MqttCmdPrefix+challenge.GetChallengee()+"/PoPInit", 0, false, blockHeight)
	token.Wait()
	err = token.Error()
	if err != nil {
		return
	}

	token = MQTTClient.Publish(MqttCmdPrefix+challenge.GetChallenger()+"/PoPInit", 0, false, blockHeight)
	token.Wait()
	err = token.Error()
	if err != nil {
		return
	}

	MQTTClient.Disconnect(1000)
	return
}

func GetMqttStatusOfParticipant(address string, responseTimeoutInMs int64) (isAvailable bool, err error) {
	LazyLoadMQTTClient()
	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}
	rwMu.RLock() // Lock for reading
	_, ok := mqttMachineByAddressAvailabilityMapping[address]
	rwMu.RUnlock() // Unlock after reading
	if ok {
		return
	}
	var messageHandler mqtt.MessageHandler = func(_ mqtt.Client, msg mqtt.Message) {
		topicParts := strings.Split(msg.Topic(), "/")
		if len(topicParts) == 3 && topicParts[1] == address {
			rwMu.Lock() // Lock for writing
			mqttMachineByAddressAvailabilityMapping[address] = true
			rwMu.Unlock() // Unlock after writing
		}
	}
	subscriptionTopic := "stat/" + address + "/STATUS"
	publishingTopic := MqttCmdPrefix + address + "/STATUS"
	// Subscribe to a topic
	if token := MQTTClient.Subscribe(subscriptionTopic, 0, messageHandler); token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}

	rwMu.Lock() // Lock for writing
	mqttMachineByAddressAvailabilityMapping[address] = false
	rwMu.Unlock() // Unlock after writing

	if token := MQTTClient.Publish(publishingTopic, 0, false, ""); token.Wait() && token.Error() != nil {
		err = token.Error()
	} else {
		duration := responseTimeoutInMs
		time.Sleep(time.Millisecond * time.Duration(duration))
	}

	// Unsubscribe and disconnect
	if token := MQTTClient.Unsubscribe(subscriptionTopic); token.Wait() && token.Error() != nil {
		err = token.Error()
	}

	rwMu.Lock() // Lock for writing
	isAvailable = mqttMachineByAddressAvailabilityMapping[address]
	delete(mqttMachineByAddressAvailabilityMapping, address)
	rwMu.Unlock() // Unlock after writing
	MQTTClient.Disconnect(1000)
	return
}

func ToJSON(payload []byte) (map[string]interface{}, error) {
	jsonString := string(payload)

	var result map[string]interface{}
	// Unmarshal the JSON string into the map
	err := json.Unmarshal([]byte(jsonString), &result)
	return result, err
}
