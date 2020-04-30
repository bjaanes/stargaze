package keeper

// TODO: Define if your module needs Parameters, if not this can be deleted

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rocket-protocol/stakebird/x/stake/types"
)

func (k Keeper) VotingPeriod(ctx sdk.Context) (res time.Duration) {
	k.paramspace.Get(ctx, types.KeyVotingPeriod, &res)
	return
}

// GetParams returns the total set of x/stake parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the x/stake parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
