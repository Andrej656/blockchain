package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/blockchain/x/registry/types"
)

func (k msgServer) AccessName(goCtx context.Context, msg *types.MsgAccessName) (*types.MsgAccessNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting name information from the store
	whois, found := k.GetWhoIs(ctx, msg.GetName())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name not found in registry")
	}

	// If the message sender address doesn't match the name owner, throw an error
	if !(msg.Creator == whois.GetCreator()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Registered name is owned by another address")
	}

	return &types.MsgAccessNameResponse{
		Did:             whois.GetDid(),
		DidDocumentJson: whois.GetDocument(),
		WhoIs:           &whois,
	}, nil
}
