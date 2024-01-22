package ante

import (
	"github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type MachineKeeper interface {
	GetMachineIndexByPubKey(ctx sdk.Context, pubKey string) (val types.MachineIndex, found bool)
	GetMachineIndexByAddress(ctx sdk.Context, address string) (val types.MachineIndex, found bool)
	GetTrustAnchor(ctx sdk.Context, pubKey string) (val types.TrustAnchor, activated bool, found bool)
}

// AccountKeeper defines the contract needed for AccountKeeper related APIs.
// Interface provides support to use non-sdk AccountKeeper for AnteHandler's decorators.
type AccountKeeper interface {
	GetParams(ctx sdk.Context) (params authtypes.Params)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	SetAccount(ctx sdk.Context, acc authtypes.AccountI)
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// FeegrantKeeper defines the expected feegrant keeper.
type FeegrantKeeper interface {
	UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error
}

type BankKeeper interface {
	IsSendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error
	SendCoins(ctx sdk.Context, from, to sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type DaoKeeper interface {
	GetMintRequestByHash(ctx sdk.Context, hash string) (val daotypes.MintRequest, found bool)
	GetMintAddress(ctx sdk.Context) (mintAddress string)
	GetClaimAddress(ctx sdk.Context) (claimAddress string)
	IsValidReissuanceProposal(ctx sdk.Context, msg *daotypes.MsgReissueRDDLProposal) (isValid bool)
	GetRedeemClaim(ctx sdk.Context, benficiary string, id uint64) (val daotypes.RedeemClaim, found bool)
}

type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
}
