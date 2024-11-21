package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	machinekeeper "github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/planetmint/planetmint-go/x/machine/simulation"
	"github.com/planetmint/planetmint-go/x/machine/testutil"
	"github.com/planetmint/planetmint-go/x/machine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"
)

type SimTestSuite struct {
	suite.Suite

	r             *rand.Rand
	accounts      []simtypes.Account
	ctx           sdk.Context
	app           *runtime.App
	bankKeeper    bankkeeper.Keeper
	accountKeeper authkeeper.AccountKeeper
	stakingKeeper *stakingkeeper.Keeper
	machineKeeper *machinekeeper.Keeper

	encCfg moduletestutil.TestEncodingConfig
}

func (s *SimTestSuite) SetupTest() {
	s.r = rand.New(rand.NewSource(1))
	accounts := simtypes.RandomAccounts(s.r, 4)

	// create genesis accounts
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	accs := []simtestutil.GenesisAccount{
		{GenesisAccount: acc, Coins: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000)))},
	}

	// create validator set with single validator
	account := accounts[0]
	tmPk, err := cryptocodec.ToTmPubKeyInterface(account.PubKey)
	require.NoError(s.T(), err)
	validator := cmttypes.NewValidator(tmPk, 1)

	startupCfg := simtestutil.DefaultStartUpConfig()
	startupCfg.GenesisAccounts = accs
	startupCfg.ValidatorSet = func() (*cmttypes.ValidatorSet, error) {
		return cmttypes.NewValidatorSet([]*cmttypes.Validator{validator}), nil
	}

	var (
		accountKeeper authkeeper.AccountKeeper
		bankKeeper    bankkeeper.Keeper
		stakingKeeper *stakingkeeper.Keeper
		machineKeeper *machinekeeper.Keeper
	)

	app, err := simtestutil.SetupWithConfiguration(testutil.AppConfig, startupCfg, &bankKeeper, &accountKeeper, &stakingKeeper, &machineKeeper)
	require.NoError(s.T(), err)

	ctx := app.BaseApp.NewContext(false, cmtproto.Header{})

	initAmt := stakingKeeper.TokensFromConsensusPower(ctx, 200)
	initCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initAmt))

	s.accounts = accounts
	// remove genesis validator account
	// add coins to the accounts
	for _, account := range accounts[1:] {
		acc := accountKeeper.NewAccountWithAddress(ctx, account.Address)
		accountKeeper.SetAccount(ctx, acc)
		s.Require().NoError(banktestutil.FundAccount(bankKeeper, ctx, account.Address, initCoins))
	}

	s.accountKeeper = accountKeeper
	s.bankKeeper = bankKeeper
	s.machineKeeper = machineKeeper
	s.ctx = ctx
	s.app = app
}

func (s *SimTestSuite) TestSimulateMsgRegisterTrustAnchor() {
	s.app.BeginBlock(abci.RequestBeginBlock{Header: cmtproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash}})

	// execute operation
	op := simulation.SimulateMsgRegisterTrustAnchor(s.accountKeeper, s.bankKeeper, *s.machineKeeper)
	operationMsg, futureOperations, err := op(s.r, s.app.BaseApp, s.ctx, s.accounts[1:], "")
	s.Require().NoError(err)

	var msg types.MsgRegisterTrustAnchor
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	s.Require().True(operationMsg.OK)
	s.Require().Equal(types.TypeMsgRegisterTrustAnchor, msg.Type())
	s.Require().Len(futureOperations, 0)
}

func TestSimTestSuite(t *testing.T) {
	suite.Run(t, new(SimTestSuite))
}
