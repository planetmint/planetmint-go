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

func SendMqttMessagesToServer(ctx sdk.Context, challenge types.Challenge) {
	err := sendMqttMessages(challenge)
	if err != nil {
		GetAppLogger().Error(ctx, "MQTT error: "+err.Error())
	}
}

func sendMqttMessages(challenge types.Challenge) (err error) {
	conf := config.GetConfig()
	hostPort := net.JoinHostPort(conf.MqttDomain, strconv.FormatInt(int64(conf.MqttPort), 10))
	uri := fmt.Sprintf("tcp://%s", hostPort)

	opts := mqtt.NewClientOptions().AddBroker(uri)
	opts.SetClientID(conf.ValidatorAddress)
	opts.SetUsername(conf.MqttUser)
	opts.SetPassword(conf.MqttPassword)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}
	blockHeight := strconv.FormatInt(challenge.GetHeight(), 10)
	token := client.Publish("cmnd/"+challenge.GetChallengee()+"/PoPInit", 0, false, blockHeight)
	token.Wait()
	err = token.Error()
	if err != nil {
		return
	}

	token = client.Publish("cmnd/"+challenge.GetChallenger()+"/PoPInit", 0, false, blockHeight)
	token.Wait()
	err = token.Error()
	if err != nil {
		return
	}

	client.Disconnect(1000)
	return
}
