package util

import (
	"fmt"
	"net"
	"strconv"

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
}

var (
	MQTTClient MQTTClientI
)

func init() {
	conf := config.GetConfig()
	hostPort := net.JoinHostPort(conf.MqttDomain, strconv.FormatInt(int64(conf.MqttPort), 10))
	uri := fmt.Sprintf("tcp://%s", hostPort)

	opts := mqtt.NewClientOptions().AddBroker(uri)
	opts.SetClientID(conf.ValidatorAddress)
	opts.SetUsername(conf.MqttUser)
	opts.SetPassword(conf.MqttPassword)
	MQTTClient = mqtt.NewClient(opts)
}

func SendMqttMessagesToServer(ctx sdk.Context, challenge types.Challenge) {
	err := sendMqttMessages(challenge)
	if err != nil {
		GetAppLogger().Error(ctx, "MQTT error: "+err.Error())
		return
	}
	GetAppLogger().Info(ctx, "MQTT message successfully sent: "+challenge.String())
}

func sendMqttMessages(challenge types.Challenge) (err error) {
	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}
	blockHeight := strconv.FormatInt(challenge.GetHeight(), 10)
	token := MQTTClient.Publish("cmnd/"+challenge.GetChallengee()+"/PoPInit", 0, false, blockHeight)
	token.Wait()
	err = token.Error()
	if err != nil {
		return
	}

	token = MQTTClient.Publish("cmnd/"+challenge.GetChallenger()+"/PoPInit", 0, false, blockHeight)
	token.Wait()
	err = token.Error()
	if err != nil {
		return
	}

	MQTTClient.Disconnect(1000)
	return
}
