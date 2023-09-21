package keeper

import (
	config "github.com/planetmint/planetmint-go/config"
	"strconv"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/hypebeast/go-osc/osc"
)

func (k Keeper) IssueResponseHandler(logger log.Logger) {
	conf := config.GetConfig()
	addr := "0.0.0.0:" + strconv.FormatInt(int64(conf.OSCServicePort), 10)
	d := osc.NewStandardDispatcher()
	err := d.AddMsgHandler("/rddl/resp", func(msg *osc.Message) {
		logger.Info("Issue Response: " + msg.String())
	})
	if err != nil {
		logger.Error("Unable to add handler to OSC service.")
	}
	server := &osc.Server{
		Addr:       addr,
		Dispatcher: d,
	}
	err = server.ListenAndServe()
	if err != nil {
		logger.Error("Unable to start the OSC service.")
	}
}
