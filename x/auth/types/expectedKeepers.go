package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

type PosKeeper interface {
	GetMsgStakeOutputSigner(sdk.Ctx, sdk.Msg) sdk.Address
}
