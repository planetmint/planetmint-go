package dao

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/lib"
	e2etestutil "github.com/planetmint/planetmint-go/testutil/e2e"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
	"github.com/stretchr/testify/suite"
)

var msgs = []sdk.Msg{
	&daotypes.MsgInitPop{},
	&daotypes.MsgDistributionRequest{},
	&daotypes.MsgDistributionResult{},
	&daotypes.MsgReissueRDDLProposal{},
	&daotypes.MsgReissueRDDLResult{},
	&machinetypes.MsgRegisterTrustAnchor{},
	&machinetypes.MsgNotarizeLiquidAsset{},
}

type RestrictedMsgsE2ESuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewRestrictedMsgsE2ESuite(cfg network.Config) *RestrictedMsgsE2ESuite {
	return &RestrictedMsgsE2ESuite{cfg: cfg}
}

func (s *RestrictedMsgsE2ESuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")
	conf := config.GetConfig()
	conf.FeeDenom = sample.FeeDenom

	s.network = network.New(s.T(), s.cfg)
	account, err := e2etestutil.CreateAccount(s.network, sample.Name, sample.Mnemonic)
	s.Require().NoError(err)
	err = e2etestutil.FundAccount(s.network, account)
	s.Require().NoError(err)
}

func (s *RestrictedMsgsE2ESuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

func (s *RestrictedMsgsE2ESuite) TestRestrictedMsgsValidator() {
	val := s.network.Validators[0]

	msg := daotypes.NewMsgInitPop(val.Address.String(), val.Address.String(), val.Address.String(), val.Address.String(), 0)
	out, err := lib.BroadcastTxWithFileLock(val.Address, msg)
	s.Require().NoError(err)

	txResponse, err := lib.GetTxResponseFromOut(out)
	s.Require().NoError(err)
	s.Require().Equal(int(0), int(txResponse.Code))
}

func (s *RestrictedMsgsE2ESuite) TestRestrictedMsgsNonValidator() {
	val := s.network.Validators[0]

	k, err := val.ClientCtx.Keyring.Key(sample.Name)
	s.Require().NoError(err)
	addr, _ := k.GetAddress()

	for _, msg := range msgs {
		msg = setCreator(msg, addr.String())
		out, err := lib.BroadcastTxWithFileLock(addr, msg)
		s.Require().NoError(err)

		txResponse, err := lib.GetTxResponseFromOut(out)
		s.Require().NoError(err)
		s.Require().Equal(int(18), int(txResponse.Code))
		s.Require().NoError(s.network.WaitForNextBlock())
	}
}

func setCreator(msg sdk.Msg, creator string) sdk.Msg {
	switch sdk.MsgTypeURL(msg) {
	case "/planetmintgo.dao.MsgInitPop":
		msg, ok := msg.(*daotypes.MsgInitPop)
		if ok {
			msg.Creator = creator
		}
	case "/planetmintgo.dao.MsgDistributionRequest":
		msg, ok := msg.(*daotypes.MsgDistributionRequest)
		if ok {
			msg.Creator = creator
		}
	case "/planetmintgo.dao.MsgDistributionResult":
		msg, ok := msg.(*daotypes.MsgDistributionResult)
		if ok {
			msg.Creator = creator
		}
	case "/planetmintgo.dao.MsgReissueRDDLProposal":
		msg, ok := msg.(*daotypes.MsgReissueRDDLProposal)
		if ok {
			msg.Creator = creator
		}
	case "/planetmintgo.dao.MsgReissueRDDLResult":
		msg, ok := msg.(*daotypes.MsgReissueRDDLResult)
		if ok {
			msg.Creator = creator
		}
	case "/planetmintgo.machine.MsgNotarizeLiquidAsset":
		msg, ok := msg.(*machinetypes.MsgNotarizeLiquidAsset)
		if ok {
			msg.Creator = creator
		}
	case "/planetmintgo.machine.MsgRegisterTrustAnchor":
		msg, ok := msg.(*machinetypes.MsgRegisterTrustAnchor)
		if ok {
			msg.Creator = creator
		}
	}
	return msg
}
