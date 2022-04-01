package keeper

import (
	"context"
	"fmt"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/blockchain/x/registry/types"
	"github.com/sonr-io/core/did"
	ssi "github.com/sonr-io/core/did/ssi"
)

// RegisterName registers a name with the registry for the given validated
func (k msgServer) RegisterName(goCtx context.Context, msg *types.MsgRegisterName) (*types.MsgRegisterNameResponse, error) {
	// Get the sender address and Generate BaseID
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting name information from the store
	whois, found := k.GetWhoIs(ctx, msg.GetNameToRegister())
	if found {
		// If the message sender address doesn't match the name owner, throw an error
		if !(msg.Creator == whois.GetCreator()) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Registered name is owned by another address")
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name already registered to this Account")
		}
	}

	// Create a new DID Document
	doc, err := GenerateDid(msg.GetCreator(), msg.GetNameToRegister(), msg.GetCredential())
	if err != nil {
		return nil, err
	}

	// Marshal the DID Document
	didJson, err := doc.MarshalJSON()
	if err != nil {
		return nil, err
	}

	// Create a new who is record
	newWhois := types.WhoIs{
		Name:     msg.GetNameToRegister(),
		Did:      doc.ID.ID,
		Document: didJson,
		Creator:  msg.GetCreator(),
	}

	newWhois.AddCredential(msg.GetCredential())

	// Write whois information to the store
	k.SetWhoIs(ctx, newWhois)

	return &types.MsgRegisterNameResponse{
		IsSuccess:       true,
		DidUrl:          doc.ID.ID,
		DidDocumentJson: didJson,
	}, nil
}

// GenerateDid generates a new did document
func GenerateDid(accountAddr, nameToRegister string, cred *types.Credential) (*did.Document, error) {
	// Generate a new DID String
	id, err := did.ParseDID(fmt.Sprintf("did:sonr:%s", accountAddr))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Get PubKey from Credential
	pubKey, err := cred.DecodePublicKey()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	verf, err := did.NewVerificationMethod(*id, ssi.JsonWebKey2020, *id, pubKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Empty did document:
	doc := &did.Document{
		Context: []ssi.URI{did.DIDContextV1URI()},
		ID:      *id,
		AlsoKnownAs: []string{
			nameToRegister,
		},
	}

	// This adds the method to the VerificationMethod list and stores a reference to the assertion list
	doc.AddAssertionMethod(verf)
	return doc, nil
}
